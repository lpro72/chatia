package cell

/*******************
* Interface
*******************/
type CellInterface interface {
    DumpCell(cell *CellStruct, indentation []byte)
}

/*******************
* Types
*******************/
type CellStruct struct {
    Count int
    NextCell *CellStruct
    PreviousCell *CellStruct
    ChildCell *CellStruct
    ParentCell *CellStruct
    Cell CellInterface
}

