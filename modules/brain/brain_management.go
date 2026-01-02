package brain

/*******************
* Import
*******************/
import (
	"chatia/modules/data"
	"chatia/modules/interfaces"
)

/*******************
* BrainManagement_Register
*******************/
func BrainManagement_Register() {
	data.BrainContextManagement_RegisterNewContext("__Management__", BrainManagement_ManagementFactory)
}

/*******************
* BrainManagement_ManagementFactory
*******************/
func BrainManagement_ManagementFactory(brainContext interfaces.I_BrainContext) {
	brainContext.SetDumpMemoryFunction(ManagementDumpMemory)
}

/*******************
* ManagementDumpMemory
*******************/
func ManagementDumpMemory(brainContext interfaces.I_BrainContext) {
	println("brain_management/ManagementDumpMemory")
	// 	brainConfig := brainContext.GetBrainConfig()
	// 	brainCellManagement := brainConfig.GetCellGroupManagement()
	// 	fmt.Printf("Number of group %d\n", brainCellManagement.GetCellGroupsCount())
	// 	for i := 0; i < brainCellManagement.GetCellGroupsCount(); i++ {
	// 		fmt.Printf("Group id %d\nCell count %d\n", i, brainCellManagement.GetCellCount(i))
	// 	}

	//	for i := 1; i <= brainCellManagement.GetCellCount(-1); i++ {
	//		cell := brainCellManagement.GetCellFromID(i)
	//		fmt.Printf("Cell id %v - ", cell)
	//		fmt.Printf("%s(%p) - %v\n", data.CellType_GetTypeName(cell.GetCellType()), cell, cell)
	//	}
}
