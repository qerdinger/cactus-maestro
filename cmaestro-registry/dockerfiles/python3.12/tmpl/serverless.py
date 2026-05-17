from cactuskit import ApiMethod, ApiProtocol, HttpStatus, cactuize

def authenticate():
    return True

@cactuize()
def simple_entrypoint():
    return f"Hello World from {simple_entrypoint}"