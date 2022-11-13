public class Utils {

  // Convert a sha256 byte array to hex String
  public static String bytesToHex(byte[] hash) {
    StringBuffer hexString = new StringBuffer();
    for (int i = 0; i < hash.length; i++) {
      String hex = Integer.toHexString(0xff & hash[i]);
      if (hex.length() == 1) {
        hexString.append('0');
      }
      hexString.append(hex);
    }
    return hexString.toString();
  }

  // Returns a reverse order of the given byte array
  public static byte[] reverseBytes(byte[] bytes) {
    byte[] buffer = new byte[bytes.length];
    for (int i = 0; i < bytes.length; i++) {
      buffer[i] = bytes[bytes.length - 1 - i];
    }
    return buffer;
  }
}



