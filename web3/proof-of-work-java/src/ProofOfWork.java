import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * Proof of work algorithm
 * @author Marcelo Soares
 *
 */
public class ProofOfWork {

  public static void main(String[] args) throws NoSuchAlgorithmException {

    //Set the difficulty by typing the amount of zeros that you want to the hash to begin
    String dificulty = args[0];

    //See more https://en.bitcoin.it/wiki/Block_hashing_algorithm

    //Block header

    //Used to keep track of protocol upgrades
    long version = 2;
    String prev_block = args[1];
    String mrkl_root = args[2];
    long timestamp = Long.parseLong(args[3]);
    //Target. Changes every 2016 blocks
    long bits = 419520339;

    long nonce = 0;

    String message = version + new String(Utils.reverseBytes(prev_block.getBytes())) + new String(Utils.reverseBytes(mrkl_root.getBytes())) + timestamp + bits;
    String hashTest = null;

    while (true) {

      MessageDigest digest = MessageDigest.getInstance("SHA-256");
      byte[] hash = digest.digest(digest.digest(message.concat(Long.toString(nonce)).getBytes(StandardCharsets.UTF_8)));
      hashTest = Utils.bytesToHex(Utils.reverseBytes(hash));
      System.out.println("Hash: " + hashTest);

      if(hashTest.substring(0, dificulty.length()).equals(dificulty))
        break;
      nonce++;

    }

    System.out.println("------------------------");
    System.out.println("Block mined!");
    System.out.println("Nonce: " + nonce);
    System.out.println("Block hash: " + hashTest);

  }

}
