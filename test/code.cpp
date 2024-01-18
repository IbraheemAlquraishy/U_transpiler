#include <iostream>

	using namespace std;
	
	int main(){
		int x; 
int y; 
cout<<"enter x:" <<endl; 
cin>>x; 
cout<<"enter y:" <<endl; 
cin>>y; 
cout<<"enter operation:" <<endl; 
char op; 
cin>>op; 
if(op == '+'){ cout<<x+y <<endl; 
} 
else if(op == '-'){ cout<<x-y <<endl; 
} 
else if(op == '*'){ cout<<x*y <<endl; 
} 
else if(op=='/'){ if(y==0){ cout<<"err" <<endl; 
} 
else{ cout<<x/y <<endl; 
} 
} 
else{ cout<<"no such operation"; 
} 

}