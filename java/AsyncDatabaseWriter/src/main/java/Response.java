import javax.xml.bind.DatatypeConverter;
import java.security.NoSuchAlgorithmException;
import java.sql.Timestamp;
import java.security.MessageDigest;

public class Response {
    String id;
    String response;
    int time_bucket;
    public static final int BUCKET_SIZE = 900;

    public Response(String r) {
        Timestamp timestamp = new Timestamp(System.currentTimeMillis());
        time_bucket = (int) (timestamp.getTime() / BUCKET_SIZE);

        MessageDigest md = null;
        try {
            md = MessageDigest.getInstance("MD5");
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }

        String raw = timestamp + r;
        md.update(raw.getBytes());
        byte[] digest = md.digest();
        id = DatatypeConverter.printHexBinary(digest).toUpperCase();

        response = r;
    }
}
