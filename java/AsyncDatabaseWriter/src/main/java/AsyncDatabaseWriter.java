import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.Session;

import java.util.concurrent.LinkedBlockingQueue;

public class AsyncDatabaseWriter implements Runnable {
    private LinkedBlockingQueue<String> queue = new LinkedBlockingQueue<String>();
    private boolean doStop = false;

    public synchronized void doStop() {
        this.doStop = true;
    }

    public void ScheduleWrite(String data) {
        queue.add(data);
    }

    @Override
    public void run() {
        // connect to database
        String serverIP = "127.0.0.1";
        String keyspace = "smapp_routing_l1_keyspace";
        String table = "peyda_output";

        Cluster cluster = Cluster.builder()
                .addContactPoints(serverIP)
                .build();

        Session session = cluster.connect(keyspace);
        System.out.println("connected");

        // you are now connected to the cluster, congrats!

        while (!doStop) {
            String data = null;
            try {
                data = queue.take();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println(data);

            // write to database
            String cqlStatementC = String.format("INSERT INTO %s.%s (output) VALUES ('%s')", keyspace, table, data);
            System.out.println(cqlStatementC);
            session.execute(cqlStatementC);
        }
    }
}

