/* Given the number of the steps of a stair you have to write a function which returns the number of ways you can go from the bottom to the top
You can only take a number of steps based on a set given to the function.


EXAMPLE:
With N=2 and X={1,2} you can go to the top using 2 steps or you can use a single step (valued 2) so return 2

Easy example with constant X
def numWays(n):
    if n== 0 or n == 1:
       return 1
    else:
       return numWays(n-1) + numWays(n-2)

We can notify that the result of the N case is the result of the sum of the 2 previosly results

N  output
0   1
1   1
2   2
3   3
4   5
...

So we can just remember the previously 2 results and sum them
*/

using namespace std;

int numWays(int n, int* x) {
	if(n == 0 || n == 1) {
		return 1;
	}
	int* nums = new int[n+1];
	nums[0] = 1;

	for(int i = 0; i < n; i++) {
		int total = 0;
		for(int j = 0; j < sizeof(x); j++) {
			total += nums[i-x[j]];
		}
		nums[i] = total;
	}
	return nums[n];
}

// Asked by Amazon
