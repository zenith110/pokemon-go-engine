from PIL import Image, Image
import os
from pathlib import Path
import imageio
import shutil
import glob
def transparent_sprite(folder_name: str) -> None:
    # Loops through the entire folder
    folder_directory = os.listdir(folder_name)
    for files in folder_directory:
        image = Image.open(f"{folder_name}/{files}")
        rgba = image.convert("RGBA")
        datas = rgba.getdata()

        most_frequent_color = max(image.getcolors(image.size[0]*image.size[1]))
        try:
            newData = [
            most_frequent_color[1][0],
            most_frequent_color[1][1],
            most_frequent_color[1][2]
            ]
            
            imgData = []

            for item in datas:
                if item[0] == newData[0] and item[1] == newData[1] and item[2] == newData[2]:
                    imgData.append((newData[0], newData[1], newData[2], 0))
                else:
                    imgData.append(item)

            rgba.putdata(imgData)
            rgba.save(f"{folder_name}/{files}", "png")
        except Exception as e:
            print(f"Skipping {files}")
            print(e)
folder_directory = os.listdir("icons")

def make_gif(frame_folder, pokemon_number):
    frames = [Image.open(image) for image in glob.glob(f"{frame_folder}/*.png")]
    frame_one = frames[0]
    frame_one.save(f"{frame_folder}/{pokemon_number}.gif", format="GIF", append_images=frames,
               save_all=True, duration=160, loop=0)

for files in folder_directory:
    images = []
    file = files.split(".")
    file_name = file[0]
    
    image = Image.open(f"icons/{files}")
    frontFolder = Path(os.getcwd() + f"/icons/{file_name}").mkdir(parents= True, exist_ok= True)
    width, height = image.size
    frame1Left = 0
    frame1Top = 0
    frame1Right = 32
    frame1Bottom = 32

    frame2Left = 0
    frame2Top = 32
    frame2Right = 64
    frame2Bottom = 64

    croppedFrame1 = image.crop((frame1Left, frame1Top, frame1Right, frame1Bottom))
    croppedFrame1.save(f"icons/{file_name}/{file_name}_frame1.png")
    croppedFrame2 = image.crop((frame2Left, frame2Top, frame2Right, frame2Bottom))
    frame2 = Image.new("RGB", (32,32), (255, 0, 0, 0))
    frame2.paste(croppedFrame2, (0,0), mask=0)
    frame2.save(f"icons/{file_name}/{file_name}_frame2.png")
    transparent_sprite(f"icons/{file_name}")
    make_gif(f"icons/{file_name}", f"{file_name}")
    print("Done")