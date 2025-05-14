def read_file(file_path):
    with open(file_path, 'r') as f:
        return f.read()
        
def generate_slug(id: int, name: str) -> str:
    # Map of Turkish characters to ASCII equivalents
    turkish_char_map = {
        'ç': 'c', 'Ç': 'C',
        'ö': 'o', 'Ö': 'O',
        'ş': 's', 'Ş': 'S',
        'ı': 'i', 'İ': 'I',
        'ğ': 'g', 'Ğ': 'G',
        'ü': 'u', 'Ü': 'U'
    }

    # Replace Turkish characters
    normalized_name = name
    for turkish, ascii in turkish_char_map.items():
        normalized_name = normalized_name.replace(turkish, ascii)

    # Format name: lowercase, trim, replace spaces with hyphens, remove non-alphanumeric
    formatted_name = normalized_name.lower().strip()
    formatted_name = '-'.join(formatted_name.split())  # Replace spaces with hyphens
    formatted_name = ''.join(c for c in formatted_name if c.isalnum() or c == '-')

    # Return slug with ID appended
    return f"{formatted_name}-{id}"