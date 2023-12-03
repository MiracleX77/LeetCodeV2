package ProgramingSkill;

public class S58 {
    public static void main(String[] args) {
        System.out.println(lengthOfLastWord("   fly me   to   the moon  "));
    }

    public static int lengthOfLastWord(String s) {
        String[] strs = s.split(" ");
        return strs[strs.length-1].length();
    }
}
