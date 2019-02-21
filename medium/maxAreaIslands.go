func maxAreaOfIsland(grid [][]int) int {
    max :=0
    for i:=0;i<len(grid);i++{
        for j:=0; j<len(grid[i]);j++{
            if grid[i][j]==1{
                area := 0
                getIslandSize(grid,i,j,&area)
                if area > max{
                    max = area
                }
            }   
        }
    }
    return max
}

func getIslandSize(grid [][]int, x int, y int, area *int){
    *area++
    grid[x][y]=3
    if x>0 && x<len(grid) && grid[x-1][y] == 1{
        getIslandSize(grid,x-1,y,area)
    }
    if x+1<len(grid) && grid[x+1][y] == 1{
        getIslandSize(grid,x+1,y,area)
    }
    if y>0  && grid[x][y-1] == 1{
        getIslandSize(grid,x,y-1,area)
    }
    if  y+1<len(grid[x]) && grid[x][y+1] == 1{
        getIslandSize(grid,x,y+1,area)
    }
}

