package brain

/*******************
* Types
*******************/
type BrainStruct struct {
    FirstCell any
    Learn func(data []byte, firstCell any)
    Exec func(command string, extraVar ...any) string
    Unittest func()

    // Debug only
    DumpMemory func()
}

/*******************
* Globals Varables
*******************/
var g_BrainFactory = make(map[string]func (brain *BrainStruct))
var g_Brain = make(map[string]*BrainStruct)

/*******************
* AddBrainFactory
*******************/
func AddBrainFactory(name string, factory func (*BrainStruct)) {
    if g_BrainFactory[name] == nil {
        g_BrainFactory[name] = factory
    }
}

/*******************
* GetBrainContext
*******************/
func GetBrainContext(name string) *BrainStruct {
    brainContext, _ := g_Brain[name]

    return brainContext
}

/*******************
* CreateBrainContext
*******************/
func CreateBrainContext(name string) *BrainStruct {
    brainContext := GetBrainContext(name)
    
    if brainContext == nil {
        g_Brain[name] = new(BrainStruct)
        brainContext = g_Brain[name]

        factory, exists := g_BrainFactory[name]
        if exists {
            factory(brainContext)
        }
        
    }

    return brainContext
}

