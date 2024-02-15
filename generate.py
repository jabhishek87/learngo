from itertools import tee, islice, chain, izip
import os
import jinja2

SRC_DIR = 'src'

def previous_and_next(some_iterable):
    prevs, items, nexts = tee(some_iterable, 3)
    prevs = chain([None], prevs)
    nexts = chain(islice(nexts, 1, None), [None])
    return izip(prevs, items, nexts)   

def get_all_files(root_dir):
    """[summary]
    
    :param root_dir: [description]
    :type root_dir: [type]
    """
    for subdir, dirs, files in os.walk(SRC_DIR):
        for _file in files:
            #print os.path.join(subdir, file)
            filepath = subdir + os.sep + _file
            if filepath.endswith(".go"):
                yield filepath

def render_template(context):
    PATH = os.path.dirname(os.path.abspath(__file__))
    TEMPLATE_ENVIRONMENT = jinja2.Environment(
        autoescape=True, loader=jinja2.FileSystemLoader(os.path.join(PATH, 'templates')),
        trim_blocks=False)
    return TEMPLATE_ENVIRONMENT.get_template('template.md.j2').render(context)

def read_file(filepath):
    with open(filepath, 'r') as myfile:
        data =  myfile.read()
    return data

def write_file(filepath, data):
    with open(filepath, 'w') as f:
        f.write(data)

def get_markdown_filename(filepath):
    filename, ext = os.path.splitext(os.path.basename(filepath))
    return filename + '.md'

def create_md_page():
    """[summary]
    
    """
    all_src_files = get_all_files(SRC_DIR)
    for previous, filepath, nxt in previous_and_next(all_src_files):

        previous =  'README.md' if not previous else get_markdown_filename(previous)
        nxt =  'README.md' if not nxt else get_markdown_filename(nxt)
        # print(previous, filepath, nxt)
        markdown = get_markdown_filename(filepath)
        if not os.path.isfile(markdown):
            print('Generating Markdown for {}'.format(filepath))
            source = read_file(filepath)
            template_str = render_template({
                'code':source,
                'filepath': filepath,
                '_next': nxt,
                '_prev': previous,   
                'title': os.path.splitext(markdown)[0]
            })
            write_file(markdown, template_str)

create_md_page()
#get_all_files(SRC_DIR)
