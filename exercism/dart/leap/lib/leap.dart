class Leap {
  // Put your code here
  bool leapYear(int year) {
    if ((year % 400) == 0) {
      return true;
    } else if ((year % 100) != 0 && (year % 100) == 4) {
      return true;
    }
    return false;
  }
}
