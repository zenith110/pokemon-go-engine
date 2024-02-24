import os


def reorganize(foldername, wrong_starting_id, correct_starting_id, image_base):
    files  = os.listdir(foldername)
    # Loops through the entire folder
    for filename in files:
        # Checks if the file we're looking at is the one we'd want to start at to rename
        if(f"{wrong_starting_id}_{image_base}.png"in filename):
            os.rename(f"{foldername}/{wrong_starting_id}_{image_base}.png", f"{foldername}/{correct_starting_id}_{image_base}.png")
            wrong_starting_id += 1 
            correct_starting_id += 1



