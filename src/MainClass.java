import java.io.*;
import java.util.ArrayList;
import java.util.Scanner;

public class MainClass {

    public static void main(String[] args) {
        MainClass objMainClass =
                new MainClass();
        objMainClass.Starter();
    }

    private void Starter() {
        FileCreator("questions.txt");
        FileCreator("loss.txt");
        FileCreator("familiarity.txt");
        FileCreator("advertising.txt");
        FileCreator("animals.txt");
        FileCreator("announcement.txt");

        MainMenu();
    }

    private static void Exit(int varIntErrorCode) {
        System.out.println("COMPUTER: Exit from program.");
        System.exit(varIntErrorCode);
    }

    private static void FileCreator(String varStringFileName) {
        try {

            File objFile = new File("output/" + varStringFileName);

            if (!(objFile.exists())) {
                if (objFile.createNewFile()) {
                    System.out.println("COMPUTER: Was created empty file '" + varStringFileName + "'.");
                }
            }

        }
        catch (IOException e) {
            System.out.println("COMPUTER: Error. " + e.getMessage() + ". File '"
                    + varStringFileName + "' was not created.");
            Exit(0);
        }
    }

    private void MainMenu() {
        try {
            System.out.println("COMPUTER: 1 == Get answer.");
            System.out.println("COMPUTER: 2 == Set answer.");
            System.out.println("COMPUTER: 0 == Close program.");

            Scanner objScanner =
                    new Scanner(System.in);
            System.out.print("USER: ");
            String varStringInput = objScanner.nextLine();

            String[] arrayStringNamesOfFiles = {
                    "questions.txt", "loss.txt", "familiarity.txt",
                    "advertising.txt", "animals.txt", "announcement.txt"
            };

            if (varStringInput.equals("0")) {
                Exit(0);
            } else {
                if (varStringInput.equals("1")) {
                    GetAnswerMenu(arrayStringNamesOfFiles);
                } else {
                    if (varStringInput.equals("2")) {
                        SetAnswerMenu(arrayStringNamesOfFiles);
                    } else {
                        System.out.println("COMPUTER: Unknown operation.");
                        MainMenu();
                    }
                }
            }
        }
        catch (Exception e) {
            System.out.println("COMPUTER: Error. " + e.getMessage() + ".");
        }

    }

    private void GetAnswerMenu(String[] arrayStringNamesOfFiles) {
        System.out.println("COMPUTER: 1 == Get 'questions'.");
        System.out.println("COMPUTER: 2 == Get 'loss'.");
        System.out.println("COMPUTER: 3 == Get 'familiarity'.");
        System.out.println("COMPUTER: 4 == Get 'advertising'.");
        System.out.println("COMPUTER: 5 == Get 'animals'.");
        System.out.println("COMPUTER: 6 == Get 'announcement'.");
        System.out.println("COMPUTER: 0 == Step back.");

        Scanner objScanner =
                new Scanner(System.in);
        System.out.print("USER: ");
        String varStringInput = objScanner.nextLine();

        if (varStringInput.equals("0")) {
            MainMenu();
        } else {
            if (varStringInput.equals("1")) {
                getAnswerFromFile(arrayStringNamesOfFiles[0]);
            } else {
                if (varStringInput.equals("2")) {
                    getAnswerFromFile(arrayStringNamesOfFiles[1]);
                } else {
                    if (varStringInput.equals("3")) {
                        getAnswerFromFile(arrayStringNamesOfFiles[2]);
                    } else {
                        if (varStringInput.equals("4")) {
                            getAnswerFromFile(arrayStringNamesOfFiles[3]);
                        } else {
                            if (varStringInput.equals("5")) {
                                getAnswerFromFile(arrayStringNamesOfFiles[4]);
                            } else {
                                if (varStringInput.equals("6")) {
                                    getAnswerFromFile(arrayStringNamesOfFiles[5]);
                                } else {
                                    System.out.println("COMPUTER: Unknown operation.");
                                    GetAnswerMenu(arrayStringNamesOfFiles);
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    private void SetAnswerMenu(String[] arrayStringNamesOfFiles) {
        System.out.println("COMPUTER: 1 == Set 'questions'.");
        System.out.println("COMPUTER: 2 == Set 'loss'.");
        System.out.println("COMPUTER: 3 == Set 'familiarity'.");
        System.out.println("COMPUTER: 4 == Set 'advertising'.");
        System.out.println("COMPUTER: 5 == Set 'animals'.");
        System.out.println("COMPUTER: 6 == Set 'announcement'.");
        System.out.println("COMPUTER: 0 == Step back.");

        Scanner objScanner =
                new Scanner(System.in);
        System.out.print("USER: ");
        String varStringInput = objScanner.nextLine();

        if (varStringInput.equals("0")) {
            MainMenu();
        } else {
            if (varStringInput.equals("1")) {
                setAnswerToFile(arrayStringNamesOfFiles[0]);
            } else {
                if (varStringInput.equals("2")) {
                    setAnswerToFile(arrayStringNamesOfFiles[1]);
                } else {
                    if (varStringInput.equals("3")) {
                        setAnswerToFile(arrayStringNamesOfFiles[2]);
                    } else {
                        if (varStringInput.equals("4")) {
                            setAnswerToFile(arrayStringNamesOfFiles[3]);
                        } else {
                            if (varStringInput.equals("5")) {
                                setAnswerToFile(arrayStringNamesOfFiles[4]);
                            } else {
                                if (varStringInput.equals("6")) {
                                    setAnswerToFile(arrayStringNamesOfFiles[5]);
                                } else {
                                    System.out.println("COMPUTER: Unknown operation.");
                                    SetAnswerMenu(arrayStringNamesOfFiles);
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    private void getAnswerFromFile(String varStringFileName) {
        try {
            TextTransfer objTextTransfer =
                    new TextTransfer();

            Scanner objScanner =
                    new Scanner(new File("output/" + varStringFileName));

            ArrayList<String> objArrayList =
                    new ArrayList<>();

            while (objScanner.hasNextLine()) {
                objArrayList.add(objScanner.nextLine() + "\n");
            }

            String varStringLineFromFile = "";

            for (String varString : objArrayList) {
                varStringLineFromFile += varString;
            }

            objTextTransfer.setData(varStringLineFromFile);

            System.out.println("COMPUTER: Success! Text was copied into clipboard.");
        }
        catch (Exception e) {
            System.out.println("COMPUTER: Error. " + e.getMessage() + ".");
        }

        MainMenu();
    }

    private void setAnswerToFile(String varStringFileName) {
        try {
            TextTransfer objTextTransfer =
                    new TextTransfer();
            String varStringLineFromFile = objTextTransfer.getData();

            OutputStreamWriter objWriter = new OutputStreamWriter(
                    new FileOutputStream("output/" + varStringFileName),"UTF-8");
            objWriter.write(varStringLineFromFile);
            objWriter.close();

            System.out.println("COMPUTER: Success! Text from clipboard was copied into file " +
                    "'" + varStringFileName + "'.");
        }
        catch (Exception e) {
            System.out.println("COMPUTER: Error. " + e.getMessage() + ".");
        }

        MainMenu();
    }

}
