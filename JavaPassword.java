public class Login {
    private static final String USERNAME = "admin";
    private static final String PASSWORD = "secret123";
    private static final String PASSWORD3 = "123454";
private static final String PASSWORD2 = "secret123444";
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.print("Enter username: ");
        String user = input.nextLine();
        System.out.print("Enter password: ");
        String passwd = input.nextLine();
        if (user.equals(USERNAME) && passwd.equals(PASSWORD)) {
            System.out.println("Access granted");
        } else {
            System.out.println("Access denied");
        }
    }
}
