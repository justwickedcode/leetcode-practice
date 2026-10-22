// 3942. Minimum Operations to Sort a Permutation (Medium)
// https://leetcode.com/problems/minimum-operations-to-sort-a-permutation/
// Runtime: 4 ms (beats 0.0%) · Memory: 16.3 MB (beats 0.0%)

func minOperations(nums []int) int {
    n := len(nums);
    g := 0; //Counter of elements greater than next element
    l := 0;//Counter of elements lesser than next element
    ig := 0; //Index of the latest element which is greater than next
    il := 0; //Index of the latest element which is lesser than next

    // We track ig and il, because in the case that g or l is only 1, it indicates that array has soluion and we need to return non -1 value, and need index of element which we need to put in right index for array to be sorted

    if n == 1{
        return 0;
    }

    //Traverse and check if array is sorted or not
    for i := 0;i < n;i++{
        if nums[i] > nums[(i+1)%n]{
            g++;
            ig = i;
        }else{
            l++;
            il = i;
        }

    }

    //if more than one is there for g or l, it indicates array isnt strictly increasing or decreasing
    if g != 1 && l!=1{
        return -1
    }

    
    if g == 1{
        if(ig == n-1){ //array is already sorted
            return 0;
        }
        if ig < n/2{ //This implies array needs to simply be moved left and nothing else
            return ig+1;
        }else{ // We now need to check if its cheaper to reverse array, move it left, then reverse it back or simply reverese it left. This comes into play for larger sized arrays only
            return min(ig+1,n-ig+1);
        }
    }
    if l == 1{
        if il < n/2{ //need to move it left as well as reverse array so extra cost compared to other case 
            return il + 2;
        }else{ 
            return n-il;
        }
    }
    return -1;
}
