public class problem_001 {
  static int multiple(int m) {
    int sum = 0;
    for (int i = 1; i < m; i++) {
      if((i % 3) == 0 || (i % 5) == 0)
        sum += i;
    }
    return sum;
  }

  public static void main(String [] args) {
    System.out.println(multiple(10));
  }
}
