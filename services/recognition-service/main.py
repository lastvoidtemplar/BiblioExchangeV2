from flask import Flask, request, jsonify
from pyzbar.pyzbar import decode
from PIL import Image

app = Flask(__name__)

@app.route('/read_barcode', methods=['POST'])
def read_barcode():
    # Check if the request contains a file
    if 'image' not in request.files:
        return jsonify({'error': 'No file part'})

    file = request.files['image']

    # Check if the file is empty
    if file.filename == '':
        return jsonify({'error': 'No selected file'})

    # Check if the file is an image
    if file and allowed_file(file.filename):
        try:
            image = Image.open(file)
            decoded_objects = decode(image)

            if decoded_objects:
                barcode_data = decoded_objects[0].data.decode('utf-8')
                return jsonify({'barcode': barcode_data})
            else:
                return jsonify({'error': 'No barcode found'})
        except Exception as e:
            return jsonify({'error': str(e)})

    return jsonify({'error': 'Unsupported file format'})

def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in {'png', 'jpg', 'jpeg', 'gif'}

if __name__ == '__main__':
    app.run(port=4000, debug=True)
