import com.datastax.oss.driver.api.core.CqlSession;
import com.datastax.oss.driver.api.core.cql.*;
import com.datastax.oss.driver.api.core.metadata.Metadata;

import java.net.InetSocketAddress;
import java.sql.Timestamp;
import java.util.concurrent.LinkedBlockingQueue;

public class AsyncDatabaseWriter implements Runnable {
    private static final int BUCKET_SIZE = 900;
    private LinkedBlockingQueue<String> queue = new LinkedBlockingQueue<String>();
    private boolean doStop = false;
    private String serverIP = "127.0.0.1";
    private String keyspace = "smapp_routing_l1_keyspace";
    private String table = "peyda_response";
//    private String datacenter = System.getenv("DB_LOCAL_DATACENTER");

    public synchronized void doStop() {
        this.doStop = true;
    }

    public void ScheduleWrite(String data) {
        queue.add(data);
    }

    @Override
    public void run() {
        // connect to database
        CqlSession session = CqlSession.builder()
                .withLocalDatacenter("teh1_os")
                .addContactPoint(new InetSocketAddress(serverIP, 9042))
                .withKeyspace(keyspace)
                .build();

        System.out.println("session string");
        System.out.println(session);

        Metadata metadata = session.getMetadata();
        System.out.println("Cluster name");
        System.out.println(metadata.getClusterName());

        System.out.println("nodes");
        metadata.getNodes().forEach(((uuid, node) -> {
            System.out.println(uuid);
            System.out.println(node);
        }));

        System.out.println("Connected to cassandra");

        String insertStatement = "INSERT INTO " + table + " (id, time_bucket, response) VALUES (?, ?, ?)";
        PreparedStatement preparedInsert = session.prepare(insertStatement);

        System.out.println("prepare statement");

        Timestamp timestamp = new Timestamp(System.currentTimeMillis());

        int batchSize = 1;
        while (!doStop & queue.size() >= batchSize) {
            BatchStatementBuilder batch = BatchStatement.builder(DefaultBatchType.LOGGED);

            try {
                long startBucket = timestamp.getTime() / BUCKET_SIZE;

                for (int i = 0; i < batchSize; i++) {
                    String data = queue.take();
                    Response response = new Response(data);

                    // Batch write to cassandra when inserts have different partition keys causes poor performance
                    if (response.time_bucket != startBucket) {
                        ScheduleWrite(data);
                        break;
                    }

                    batch.addStatement(preparedInsert.bind(response.id, response.time_bucket, response.response));
                }

            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            // write to database
            session.execute(batch.build());

        }
    }
}
