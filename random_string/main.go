package main

import (
	"bufio"
	"fmt"
	"github.com/pochard/commons/randstr"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var fileName string = "./tmp/dic4.txt"

//var str string = "↕↖↗↘↙↚↛↜↝↞↟↠↡↢↣↤↥↦↧↨↩↪↫↬↭↮↯↰↱↲↳↴↶↷↸↹↺↻↼↽↾↿⇀⇁⇂⇃⇄⇅⇆⇇⇈⇉⇊⇋⇌⇍⇎⇏⇕⇖⇗⇘⇙⇚⇛⇜⇝⇞⇟⇠⇡⇢⇣⇤⇥⇦⇧⇨⇩⇪⌅⌆⌤⏎▶☇☈☊☋☌☍➔➘➙➚➛➜➝➞➟➠➡➢➣➤➥➦➧➨➩➪➫➬➭➮➯➱➲➳➴➵➶➷➸➹➺➻➼➽➾⤴⤵↵↓↔←→↑༺࿈༻、❀༒❀、༺༽༾ཊ࿈ཏ༿༼༻、༺❀ൢ༒ൢ❀༻、༺ཌ༈༒༈ད༻、༺༒༻、༺࿈༻、❀༒❀、༺༽༾ཊ࿈ཏ༿༼༻、༺❀ൢ༒ൢ❀༻、༺ཌ༈༒༈ད༻ˍ∎⊞⊟⊠⊡⋄⎔▀▁▂▃▄▅▆▇█▉▊▋▋▌▍▎▏▐░▒▓▔■□▢▣▤▥▦▧▨▩▪▫▬▭▮▯▰▱►◄◆◇◈◘◙◚◛◢◣◤◥◧◨◩◪◫☖☗❏❐❑❒❖❘❙❚◊`ˊᐟ‐‑‒―⁃≣⋐⋑⋒⋓⌒⌜⌝⌞⌟⎯─━│┃┄┅┆┇┈┉┊┋┌┍┎┏┐┑┒┓└└┕┖┗┘┙┚┛├├┝┞┟┠┡┢┣┤┥┦┧┨┩┪┫┬┭┮┯┰┱┲┳┴┵┶┷┸┹┺┻┼┽┾┿╀╁╂╃╄╅╆╇╈╉╊╋╌╍╎╏══║╒╓╔╔╔╕╕╖╖╗╗╘╙╚╚╛╛╜╜╝╝╞╟╟╠╡╡╢╢╣╣╤╤╥╥╦╦╧╧╨╨╩╩╪╪╫╬╬╭╮╯╰╱╲╳╴╵╶╷╸╹╺╻╼╽╾╿▏▕◜◝◞◟◠◡☰☱☲☳☴☵☶☷✕≡⌈⌊—⌉⌋¹²³⁰ⁱ⁴⁵⁶⁷⁸⁹⁺⁻⁼⁽⁾ⁿ₀₁₂₃₄₅₆₇₈₉₊₋₌₍₎ₐₑₒₓₔぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをんゔゕゖ゚゛゜ゝゞゟ゠ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶヷヸヹヺ・ーヽヾヿ㍐㍿αβχδεηγικλμνωοφπψρστθυξζ"
var str string = "ｦｧｨｩｪｫｬｭｮｯｰｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝ"
var maxLine = 100

func main() {

	str = strings.ReplaceAll(str, " ", "")
	for i := 105; i < 106; i++ {
		fileIndex := fmt.Sprintf("./tmp/dic%d.txt", i)
		generateTxt2(fileIndex)
	}

}

func generateTxt(fileName string) {
	newPath := filepath.Dir(fileName)
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, fs.ModePerm)
	}
	//fi, err := os.Open(fileName)
	fiW, err := os.Create(fileName)
	defer fiW.Close()
	// 创建 Reader
	w := bufio.NewWriter(fiW)
	defer w.Flush()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxLine; i++ {
		maxNum := 10 + rand.Intn(10)
		line := randstr.RandomAlphabetic(maxNum)
		w.WriteString(line)
		w.WriteString("\r\n")
		time.Sleep(time.Microsecond * 100)
	}
}

func generateTxt2(fileName string) {
	newPath := filepath.Dir(fileName)
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, fs.ModePerm)
	}
	//fi, err := os.Open(fileName)
	fiW, err := os.Create(fileName)
	defer fiW.Close()
	// 创建 Reader
	w := bufio.NewWriter(fiW)
	defer w.Flush()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxLine; i++ {
		maxNum := 10 + rand.Intn(10)
		line := randstr.Random(maxNum, str)
		w.WriteString(line)
		w.WriteString("\r\n")
		time.Sleep(time.Microsecond * 100)
	}
}
