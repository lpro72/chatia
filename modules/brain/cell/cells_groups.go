package cell

/*******************
* Import
*******************/
import (
    "fmt"
)

/*******************
* Types
*******************/
type s_CellsGroup struct {
    CellCount int
    CellList []I_CellManagement
}

type s_CellType struct {
    ID int
    Name string
}

type s_CellsGroupManagement struct {
    registeredCellTypeList map[string]s_CellType
    CellGroupList []s_CellsGroup
}

/*******************
* Globals Varables
*******************/
var g_CellsGroupManagement s_CellsGroupManagement = s_CellsGroupManagement{registeredCellTypeList: make(map[string]s_CellType, 10), CellGroupList: make([]s_CellsGroup,0)}

/*******************
* RegisterCellType
*******************/
func RegisterCellType(name string) int {
    cellTypeID := GetCellTypeID(name)
    if cellTypeID == 0 {
        cellTypeID = len(g_CellsGroupManagement.registeredCellTypeList)
        g_CellsGroupManagement.registeredCellTypeList[name] = s_CellType{ID: cellTypeID, Name: name}
    }
    
    return(cellTypeID)
}

/*******************
* GetCellTypeID
*******************/
func GetCellTypeID(name string) int {
    cellType, notExists := g_CellsGroupManagement.registeredCellTypeList[name]
    if notExists {
        return 0
    }
    
    return(cellType.ID)
}

/*******************
* GetCellTypeName
*******************/
func GetCellTypeName(cellTypeID int) string {
    for cellTypeName, cellType := range g_CellsGroupManagement.registeredCellTypeList {
        if cellType.ID == cellTypeID {
            return cellTypeName
        }
    }

    return ""
}

/*******************
* AppendCellToGroup
*******************/
func AppendCellToGroup(cell I_CellManagement) int {
    lastCellGroupID := len(g_CellsGroupManagement.CellGroupList) - 1
    if lastCellGroupID == -1 || g_CellsGroupManagement.CellGroupList[lastCellGroupID].CellCount >= 1024 {
        newCellGroupe := s_CellsGroup{CellCount: 0, CellList: make([]I_CellManagement, 0, 1024)}
        g_CellsGroupManagement.CellGroupList = append(g_CellsGroupManagement.CellGroupList, newCellGroupe)
        lastCellGroupID += 1
    }
    cellGroup := &g_CellsGroupManagement.CellGroupList[lastCellGroupID]
    cellGroup.CellList = append(cellGroup.CellList, cell)
    cellGroup.CellCount += 1

    return lastCellGroupID * 1024 + cellGroup.CellCount
}

/*******************
* GetCellFromGroup
*******************/
func GetCellFromGroup(cellID int) I_CellManagement {
    if cellID == 0 {
        return nil
    }
    
    cellID -= 1
    groupeID := cellID / 1024
    cellIDInGroup := cellID % 1024
    
    if groupeID > len(g_CellsGroupManagement.CellGroupList) - 1 {
        return nil
    }
    cellGroup := g_CellsGroupManagement.CellGroupList[groupeID]
    if cellIDInGroup >= cellGroup.CellCount {
        return nil
    }
    return cellGroup.CellList[cellIDInGroup]
}

/*******************
* ManagementDumpMemory
*******************/
func ManagementDumpMemory() {
    fmt.Println("ManagementDumpMemory")
    fmt.Printf("Number of group %d\n", len(g_CellsGroupManagement.CellGroupList))
    for _, cellGroup := range g_CellsGroupManagement.CellGroupList {
        fmt.Printf("Group size %d\nCell count %d\n", len(cellGroup.CellList), cellGroup.CellCount)
        for _, cell := range cellGroup.CellList {
            fmt.Printf("%s(%p) - %v\n", GetCellTypeName(cell.GetStruct().CellType), cell, cell)
        }
    }
}

