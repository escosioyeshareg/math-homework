import sys

def main():
    # Code to be executed when the script is run

if __name__ == "__main__":
    if len(sys.argv) > 1 and sys.argv[1] == 'test':
        print("Running test case")
    else:
        main()
