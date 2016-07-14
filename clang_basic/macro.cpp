#include <stdio.h>
#include <iomanip>
#include <stdlib.h>

using namespace std;

#define max(a, b)( (a) > (b) ? (a) : (b) )

void main()
{
	int m = 0, n = 0;
	cout << max(m, ++n) << endl;
	cout << m << setw(2) << endl;

	system("pause");
}
