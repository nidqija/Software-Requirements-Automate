import shutil
from pathlib import Path

def clean_pycache(root_dir="."):
    # Convert string path to a Path object
    root = Path(root_dir)
    
    # Search for all directories named __pycache__ recursively
    # and .pyc files
    print(f"🧹 Cleaning up {root.absolute()}...")
    
    count = 0
    for folder in root.rglob("__pycache__"):
        if folder.is_dir():
            shutil.rmtree(folder)
            print(f"Removed: {folder}")
            count += 1
            
    if count == 0:
        print("✨ Everything is already sparkling clean!")
    else:
        print(f"✅ Successfully deleted {count} __pycache__ folders.")

if __name__ == "__main__":
    clean_pycache()






    