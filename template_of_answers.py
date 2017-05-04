# coding: utf8


import gtk
import os


def starter():

    try:

        if os.path.exists("output") is False:
            os.mkdir("output")
            print("COMPUTER: Was created directory \"output\".")

        if os.path.exists("output/question.txt") is False:
            open("output/question.txt", "w")
            print("COMPUTER: Was created file \"question.txt\".")

        if os.path.exists("output/loss.txt") is False:
            open("output/loss.txt", "w")
            print("COMPUTER: Was created file \"loss.txt\".")

        if os.path.exists("output/familiarity.txt") is False:
            open("output/familiarity.txt", "w")
            print("COMPUTER: Was created file \"familiarity.txt\".")

        if os.path.exists("output/advertising.txt") is False:
            open("output/advertising.txt", "w")
            print("COMPUTER: Was created file \"advertising.txt\".")

        if os.path.exists("output/animals.txt") is False:
            open("output/animals.txt", "w")
            print("COMPUTER: Was created file \"animals.txt\".")

        if os.path.exists("output/announcement.txt") is False:
            open("output/announcement.txt", "w")
            print("COMPUTER: Was created file \"announcement.txt\".")

        if os.path.exists("output/ban.txt") is False:
            open("output/ban.txt", "w")
            print("COMPUTER: Was created file \"ban.txt\".")
    except Exception as var_except:
        print(
            "COMPUTER: Error, " + str(var_except) + ". Exit from program...")
        exit()

    main_menu()


def main_menu():
    print("\n" +
          "COMPUTER [Main menu]: You are in Main menu...")
    print("COMPUTER [Main menu -> ]: 1 == Get answer.")
    print("COMPUTER [Main menu -> ]: 2 == Set answer.")
    print("COMPUTER [Main menu -> ]: 0 == Close program.")

    user_answer = raw_input("USER [Main menu -> Set answer]: ")

    if user_answer == "0":
        exit()
    else:
        if user_answer == "1":
            get_answer()
        else:
            if user_answer == "2":
                set_answer()
            else:
                print("COMPUTER [Main menu]: Unknown command. Retry query...")
                main_menu()


def get_answer():

    def algorythm_get_answer(answer_type):
        try:
            cb = gtk.clipboard_get()

            file = open("output/" + answer_type + ".txt")
            text = file.read()
            file.close()

            cb.set_text(text)
            gtk.timeout_add(1, gtk.main_quit)
            gtk.main()

            print("COMPUTER [.. -> Set answer -> Set " + answer_type +
                  "]: Answer was successfully copied to clipboard. Return to Main menu...")
        except Exception as var_except:
            print("COMPUTER [Main Menu -> Get answer]: Error, " + str(var_except) +
                  ". Return to Main Menu...")

        main_menu()

    print("\n" +
          "COMPUTER [Main menu -> Get answer]: You are in menu Get answer...")
    print("COMPUTER [Main menu -> Get answer -> ]: 1 == Get question.")
    print("COMPUTER [Main menu -> Get answer -> ]: 2 == Get loss.")
    print("COMPUTER [Main menu -> Get answer -> ]: 3 == Get familiarity.")
    print("COMPUTER [Main menu -> Get answer -> ]: 4 == Get advertising.")
    print("COMPUTER [Main menu -> Get answer -> ]: 5 == Get animals.")
    print("COMPUTER [Main menu -> Get answer -> ]: 6 == Get announcement.")
    print("COMPUTER [Main menu -> Get answer -> ]: 7 == Get ban.")
    print("COMPUTER [Main menu -> Get answer -> ]: 0 == Step back.")

    user_answer = raw_input("USER [Main menu -> Set answer]: ")

    if user_answer == "0":
        main_menu()
    else:
        if user_answer == "1":
            algorythm_get_answer("question")
        else:
            if user_answer == "2":
                algorythm_get_answer("loss")
            else:
                if user_answer == "3":
                    algorythm_get_answer("familiarity")
                else:
                    if user_answer == "4":
                        algorythm_get_answer("advertising")
                    else:
                        if user_answer == "5":
                            algorythm_get_answer("animals")
                        else:
                            if user_answer == "6":
                                algorythm_get_answer("announcement")
                            else:
                                if user_answer == "7":
                                    algorythm_get_answer("ban")
                                else:
                                    print("COMPUTER [Main menu -> Get answer]: " +
                                          "Unknown command. Retry query...")
                                    get_answer()
    main_menu()


def set_answer():
    def algorythm_set_answer(answer_type):

        try:
            cb = gtk.clipboard_get()
            text = str(gtk.Clipboard.wait_for_text(cb))

            text = text.decode("utf8")

            file = open("output/" + answer_type + ".txt", "w")
            file.write(text)
            file.close()

            print("COMPUTER [.. -> Set answer -> Set " + answer_type +
                  "]: Answer was successfully writen. Return to Main menu...")
        except Exception as var_except:
            print("COMPUTER [Main Menu -> Set answer]: Error, " + str(var_except) +
                  ". Return to Main Menu...")
        main_menu()

    print("\n" +
          "COMPUTER [Main menu -> Set answer]: You are in menu Set answer...")
    print("COMPUTER [Main menu -> Set answer -> ]: 1 == Set question.")
    print("COMPUTER [Main menu -> Set answer -> ]: 2 == Set loss.")
    print("COMPUTER [Main menu -> Set answer -> ]: 3 == Set familiarity.")
    print("COMPUTER [Main menu -> Set answer -> ]: 4 == Set advertising.")
    print("COMPUTER [Main menu -> Set answer -> ]: 5 == Set animals.")
    print("COMPUTER [Main menu -> Set answer -> ]: 6 == Set announcement.")
    print("COMPUTER [Main menu -> Set answer -> ]: 7 == Set ban.")
    print("COMPUTER [Main menu -> Set answer -> ]: 0 == Step back.")

    user_answer = raw_input("USER [Main menu -> Set answer]: ")

    if user_answer == "0":
        main_menu()
    else:
        if user_answer == "1":
            algorythm_set_answer("question")
        else:
            if user_answer == "2":
                algorythm_set_answer("loss")
            else:
                if user_answer == "3":
                    algorythm_set_answer("familiarity")
                else:
                    if user_answer == "4":
                        algorythm_set_answer("advertising")
                    else:
                        if user_answer == "5":
                            algorythm_set_answer("animals")
                        else:
                            if user_answer == "6":
                                algorythm_set_answer("announcement")
                            else:
                                if user_answer == "7":
                                    algorythm_set_answer("ban")
                                else:
                                    print("COMPUTER [Main menu -> Set answer]: " +
                                          "Unknown command. Retry query...")
                                    set_answer()

    main_menu()


starter()
