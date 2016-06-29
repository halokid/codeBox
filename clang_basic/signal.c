#include<stdio.h>
#include<signal.h>
#include<unistd.h>
#include<stdlib.h>
void when_alarm();
void when_sigint();
void when_sigchld(int);
void when_sigusr1();
void when_sigio();
int main()
{
    int childpid;//子程序进程ID号
    printf("程序已经开始运行，5秒钟后将接收到时钟信号。\n");
    if ((childpid=fork())>0)//父进程
    {
        signal(SIGALRM,when_alarm);  //当接收到SIGALRM信号时，调用when_alarm函数
        signal(SIGINT,when_sigint);  //当接收到SIGINT信号时，调用when_sigint函数
        signal(SIGCHLD,when_sigchld);//当接收到SIGCHLD信号时，调用when_sigchld函数
        signal(SIGUSR1,when_sigusr1);//当接收到SIGUSR1信号时，调用when_sigusr1函数
        signal(SIGIO,when_sigio);//当接收到SIGIO信号时，调用when_sigio函数
        alarm(5);     //5秒钟之后产生SIGALRM信号
        raise(SIGIO); //向自己发送一个SIGIO信号
        pause(); //将父进程暂停下来，等待SIGALRM信号到来
        pause(); //将父进程暂停下来，等待SIGUSR1信号到来
        pause(); //将父进程暂停下来，等待SIGCHLD信号到来
        printf("------此时程序会停下来等待，请按下ctrl+c送出SIGINT信号-------\n");
        pause(); //将父进程暂停下来，等待SIGINT信号到来        
    }
    else if(childpid==0) //子进程
    {
        int timer;
        for(timer=7;timer>=0;timer--) //时钟计时5秒产生SIGALRM信号，再过2秒子进程退出，产生SIGCHLD信号
        {
            if(timer>2)    
                printf("距离SIGALRM信号到来还有%d秒。\n",timer-2);
            if(timer==4)
                kill(getppid(),SIGUSR1); //向父进程发送一个SIGUSR1信号
            if((timer<=2)&&(timer>0))
                printf("子进程还剩%d秒退出，届时会产生SIGCHLD信号。\n",timer);
            if(timer==0) //子进程退出，产生SIGCHLD信号
                raise(SIGKILL); //子进程给自己发一个结束信号
            sleep(1); //每个循环延时1秒钟
        }        
    }
    else
        printf("fork()函数调用出现错误！\n");
    return 0;
}
void when_alarm()
{
    printf("5秒钟时间已到，系统接收到了SIGALRM信号！\n");
}
void when_sigint()
{
    printf("已经接收到了SIGINT信号，程序将退出！\n");
    exit(0);
}
void when_sigchld(int SIGCHLD_num)
{
    printf("收到SIGCHLD信号，表明我的子进程已经中止，SIGCHLD信号的数值是：%d。\n",SIGCHLD_num);
}
void when_sigusr1()
{
    printf("系统接收到了用户自定义信号SIGUSR1。\n");
}
void when_sigio()
{
    printf("系统接收到了SIGIO信号。\n");
}
