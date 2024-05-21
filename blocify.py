from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def home():
    # Render the home page
    return render_template("index.html")

@app.route("/block/<block_hash>")
def block(block_hash):
    # Render the block details page
    # For now, just display the block hash
    return "Block Hash: " + block_hash

if __name__ == "__main__":
    app.run(debug=True)
