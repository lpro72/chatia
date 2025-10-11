package cell

/*******************
* Interface
*******************/
type I_CellData interface {
    DumpCell(cell I_CellManagement, indentation []byte)
}

type I_CellManagement interface {
    GetData () I_CellData
    GetID () int
    GetParentCell () I_CellManagement
    GetNextCell () I_CellManagement
    GetPreviousCell () I_CellManagement
    GetFirstChildCell () I_CellManagement
    GetStruct () *s_Cell
    DumpCell(indentation []byte)
}

/*******************
* Types
*******************/
type s_Cell struct {
    ID int
    CellType int
    nextCellID int
    previousCellID int
    childCellIDList []int
    parentCellID int
    CellData I_CellData
}

/*******************
* Cell management
*******************/
func (currentCell *s_Cell) GetID () int {
    return currentCell.ID
}

func (currentCell *s_Cell) GetParentCell () I_CellManagement {
    return GetCellFromGroup(currentCell.parentCellID)
}

func (currentCell *s_Cell) GetNextCell () I_CellManagement {
    return GetCellFromGroup(currentCell.nextCellID)
}

func (currentCell *s_Cell) GetPreviousCell () I_CellManagement {
    return GetCellFromGroup(currentCell.previousCellID)
}

func (currentCell *s_Cell) GetFirstChildCell () I_CellManagement {
    if currentCell.childCellIDList == nil {
        return nil
    }
    return GetCellFromGroup(currentCell.childCellIDList[0])
}

func (currentCell *s_Cell) GetData () I_CellData {
    return currentCell.CellData
}

func (currentCell *s_Cell) GetStruct () *s_Cell {
    return currentCell
}

func (currentCell *s_Cell) addChildCell(newCell *s_Cell) {
    var childCell I_CellManagement = nil
    if currentCell.childCellIDList == nil {
        currentCell.childCellIDList = make([]int,0)
    } else {
        id := len(currentCell.childCellIDList) - 1
        childCell = GetCellFromGroup(currentCell.childCellIDList[id])
    }
    if childCell != nil {
        childCell.GetStruct().nextCellID = newCell.GetID()
        newCell.previousCellID = childCell.GetID()
    }
    currentCell.childCellIDList = append(currentCell.childCellIDList,newCell.GetID())
    newCell.parentCellID = currentCell.GetID()
}

func (currentCell *s_Cell) DumpCell(indentation []byte) {
    cellData := currentCell.GetData()
    if cellData != nil {
        cellData.DumpCell(currentCell, indentation)
    }
}

/*******************
* CreateCell
*******************/
func CreateCell(parentCell I_CellManagement, currentData I_CellData, cellType int) I_CellManagement {
    newCell := new(s_Cell)
    newCell.CellType = cellType
    newCell.CellData = currentData
    newCell.ID = AppendCellToGroup(newCell)

    if parentCell != nil {
        parentCell.GetStruct().addChildCell(newCell)
    }

    return(newCell)
}


