/* Given a tree of integers and return the number of non-empty unival trees
A unival tree is a tree which contain all nodes with the same number.
*/
#include <stddef.h>
using namespace std;
struct Node{
  int value;
  Node *left;
  Node *right;
};
//DUMB SOLUTION

//O(n)
bool isUnivalDumb(Node* root) {
	if(root == NULL) {
		return true;
	}
	if(root->left != NULL && root->left->value != root->value) {
		return false;
	}
	if(root->right != NULL && root->right->value != root->value) {
		return false;
	}
	if(isUnivalDumb(root->left) && isUnivalDumb(root->right)) {
		return true;
	}
	return false;
}

// O(n^2)
int countUnivalsDumb(Node* root) {
	if(root == NULL) {
		return 0;
	}

	int count = countUnivalsDumb(root->left) + countUnivalsDumb(root->right);

	if(isUnivalDumb(root)) {
		count++;
	}
	return count;
}

// GOOD SOLUTION O(n)
struct result{
    int value;
    bool status;
};

result exploreTree(Node* root) {
	if(root == NULL){
        result r = {0, true};
		return r;
	}
	result result_left = exploreTree(root->left);
	result result_right = exploreTree(root->right);

	bool isUnival = true;
	if(!result_left.status || !result_right.status){
		isUnival = false;
	}

	if(root->left != NULL && root->left->value != root->value){
		isUnival = false;
	}

	if(root->right != NULL && root->right->value != root->value) {
		isUnival = false;
	}
    result r;
	if(isUnival) {
        r = {result_left.value + result_right.value + 1, true};
	} else {
		r = {result_left.value + result_right.value, false};
	}
    return r;
}

int countUnivals(Node* root)  {
	result r = exploreTree(root);
	return r.value;
}

// Asked by Google
