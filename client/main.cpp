#include <iostream>


#include <unistd.h>
#include <termios.h>
#include "gui.h"
Gui gui = Gui();

char getArrows();
bool run=true;

int main(int argc, char** argv) {

	gui.update(0);
	while(run)
	{
		int c = (int)getArrows();
		gui.update(c);
	}

}

char getArrows() {
    char buf = 0;
    struct termios old = {0};
    if (tcgetattr(0, &old) < 0)
            perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if (tcsetattr(0, TCSANOW, &old) < 0)
            perror("tcsetattr ICANON");
    read(0, &buf, 1);
	if(((int)buf)==27)		//if 27 is the result, read twice again to get the arrow key value.
	{
	    read(0, &buf, 1);
	    read(0, &buf, 1);
	}
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if (tcsetattr(0, TCSADRAIN, &old) < 0)
            perror ("tcsetattr ~ICANON");
    return (buf);
}


