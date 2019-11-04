#include "textflag.h"

TEXT ·cpuid(SB),NOSPLIT,$0-8
	MOVL	$1, AX
	CPUID
	MOVL	CX, a(FP)
	RET

TEXT ·rdrand(SB),NOSPLIT,$0-8
	RDRANDQ	DI
	MOVQ	DI, ret(FP)
	RET

