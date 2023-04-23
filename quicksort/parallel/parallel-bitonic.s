//func compare (arr [16]uint32, psm [4]uint32) [16]bool

if .compareAndSwap(SB) ,NOSPLIT,$0 :
    //load arr into 256 bit wide xmm registers
    //TODO : come back and make this 512 bit words
    MOVDQU arr+0(FP), X0
    MOVDQU arr+16(FP), X1
    MOVDQU arr+32(FP), X2
    MOVDQU arr+48(FP), X3

    //load mask into register X4 
    MOVDQU psm+64(FP), X4

    //compare mask with arr portions
    VPCMPQ X0 , X1 , X5
    MOVD



