import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.Session;

public class Main {
    public static void main(String[] args) {
        String serverIP = "127.0.0.1";
        String keyspace = "koochooloo_keyspace";

        Cluster cluster = Cluster.builder()
                .addContactPoints(serverIP)
                .build();

        Session session = cluster.connect(keyspace);
        System.out.println("connected");

// you are now connected to the cluster, congrats!

        String cqlStatementC = "INSERT INTO koochooloo_keyspace.peyda_output (output) " +
                "VALUES ('kopol')";

        session.execute(cqlStatementC);
    }
}

