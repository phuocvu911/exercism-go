package change
import (
    "errors"
    "math"
)
func Change(coins []int, target int) ([]int, error) {
	if target <0{
        return nil, errors.New("err")
    }
    if target == 0{
        return []int{}, nil
    }
    if len(coins) ==0{
        return nil, errors.New("coin list mt")
    }
    dp:= make([]int, target + 1)
    rc:= make([]int, target + 1)
    
    for i:=1; i<=target; i++{
        dp[i] = math.MaxInt32
        rc[i] = -1
    }
    
    for i:= 1; i<=target; i++{
        for _, c:= range coins{
            if c<=i && dp[i-c] != math.MaxInt32{
                if dp[i-c]+1 < dp[i]{
                	dp[i] = dp[i-c] +1
                	rc[i] = c
            	}
            }          
        }
    }

    if dp[target] == math.MaxInt32{
        return nil, errors.New("cant find solulu")
    }

    res:= []int{}
    for cur:= target; cur>0; cur -= rc[cur]{
        res = append(res, rc[cur])
    }
    return res, nil
}

