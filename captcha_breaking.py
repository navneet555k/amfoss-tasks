from PIL import Image
import pytesseract
pytesseract.pytesseract.tesseract_cmd = r'/usr/local/Cellar/tesseract/4.1.1/bin/tesseract'
pic=Image.open("/Users/navneet/Desktop/untitled.001.jpeg")
string=pytesseract.image_to_string(pic,lang='eng')
print(string)
