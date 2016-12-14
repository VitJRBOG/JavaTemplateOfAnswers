import java.awt.*;
import java.awt.datatransfer.*;
import java.io.IOException;

public class TextTransfer implements ClipboardOwner {
    @Override
    public void lostOwnership(Clipboard clipboard, Transferable contents) {

    }

    void setData(String data){
        StringSelection stringSelection;
        stringSelection = new StringSelection(data);
        Clipboard clipboard = Toolkit.getDefaultToolkit().getSystemClipboard();
        clipboard.setContents(stringSelection, this);
    }

    String getData() throws IOException, UnsupportedFlavorException {
        Clipboard clipboard = Toolkit.getDefaultToolkit().getSystemClipboard();
        return (String) clipboard.getData(DataFlavor.stringFlavor);
    }
}