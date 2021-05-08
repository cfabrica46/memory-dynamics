package main

//#include<stdio.h>
//#include<stdlib.h>
//int i;
//int N;
//int *vector;
//int memoria(){
//	printf("Introduzca un valor para N: ");
//	scanf("%i",&N);
//	vector = (int*)malloc(N*sizeof(int));
//	if (vector == NULL){
//		printf("Error al reservar memoira\n");
//	}else{
//		for(i = 0;i < N; i++){
//			*(vector + i) = i;
//		}
//		for(i = 0;i < N; i++){
//			printf("%i-",*(vector+i));
//		}
//		printf("\n");
//	}
//	printf("Introduzca un valor para N: ");
//	scanf("%i",&N);
//	vector = (int*)malloc(N*sizeof(int));
//	if (vector == NULL){
//		printf("Error al reservar memoira\n");
//	}else{
//		for(i = 0;i < N; i++){
//			*(vector + i) = i;
//		}
//		for(i = 0;i < N; i++){
//			printf("%i-",*(vector+i));
//		}
//		printf("\n");
//	}
//}
import "C"

func main() {

	C.memoria()

}
