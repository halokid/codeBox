import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * Proof of work verifier
 */
public class ProofOfWorkVerifier {

  public static void main(String[] args) throws NoSuchAlgorithmException {
    //Set the difficulty by typing the amount of zeros that you want to the hash to begin
    String dificulty = args[0];

    //Used to keep track of protocol upgrades
    long version = 2;
    String prev_block = args[1];
    String mrkl_root = args[2];
    long timestamp = Long.parseLong(args[3]);
    //Target. Changes every 2016 blocks
    long bits = 419520339;
    long nonce = Long.parseLong(args[4]);

    String message = version + new String(Utils.reverseBytes(prev_block.getBytes())) + new String(Utils.reverseBytes(mrkl_root.getBytes())) + timestamp + bits;

    MessageDigest digest = MessageDigest.getInstance("SHA-256");
    byte[] hash = digest.digest(digest.digest(message.concat(Long.toString(nonce)).getBytes(StandardCharsets.UTF_8)));
    String hashTest = Utils.bytesToHex(Utils.reverseBytes(hash));

    System.out.println(hashTest);

    if (hashTest.substring(0, dificulty.length()).equals(dificulty))
      System.out.println("Valid Block!");
    else
      System.out.println("Invalid!");

  }
}


