package brain

/*******************
* Import
*******************/
import (
	"fmt"

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
	brainConfig := brainContext.GetBrainConfig()
	CellGroupsManagement := brainConfig.GetCellsGroupManagament()
	SynapsesGroupsManagement := brainConfig.GetSynapsesGroupManagement()
	cellsCount := CellGroupsManagement.GetCellsCount()
	synapsesCount := SynapsesGroupsManagement.GetSynapsesCount()
	fmt.Printf("Number of group %d (Synapses)\n", SynapsesGroupsManagement.GetSynapsesGroupsCount())
	fmt.Printf("Number of synapses %d\n", synapsesCount)
	fmt.Printf("Number of group %d (Cells)\n", CellGroupsManagement.GetCellGroupsCount())
	fmt.Printf("Number of cells %d\n", cellsCount)

	for i := uint32(1); i <= synapsesCount; i++ {
		synapse := SynapsesGroupsManagement.GetSynapseFromID(i)
		fmt.Printf("Synapse id %d - %v\n", i, synapse)
	}

	for i := uint32(1); i <= cellsCount; i++ {
		cell := CellGroupsManagement.GetCellFromID(i)
		fmt.Printf("Cell id %d (%s) - %v\n", i, data.CellType_GetTypeName(cell.GetCellType()), cell)
		cell.DumpCell()
	}
}
