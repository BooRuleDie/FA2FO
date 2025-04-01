import { useState } from "react";

interface Color {
    name: string;
    value: string;
}

interface Product {
    name: string;
    description: string;
    price: number;
    colors: Color[];
    images: string[];
}

interface AppProps {
    initialProduct: Product;
}

function App({ initialProduct }: AppProps) {
    const [product] = useState<Product>(initialProduct);
    const [selectedColorIndex, setSelectedColorIndex] = useState(0);
    const [selectedImageIndex, setSelectedImageIndex] = useState(0);
    const [quantity, setQuantity] = useState(1);

    const handleColorSelect = (index: number) => {
        setSelectedColorIndex(index);
    };

    const handleImageSelect = (index: number) => {
        setSelectedImageIndex(index);
    };

    const increaseQuantity = () => {
        setQuantity((prev) => prev + 1);
    };

    const decreaseQuantity = () => {
        setQuantity((prev) => (prev > 1 ? prev - 1 : 1));
    };

    const addToCart = () => {
        alert(`Added ${quantity} of ${product.name} to cart!`);
    };

    return (
        <div className="container mx-auto px-4 py-8">
            <div className="bg-white rounded-lg shadow-lg overflow-hidden">
                <div className="md:flex">
                    <div className="md:w-1/2">
                        <div className="h-64 bg-gray-200 md:h-full">
                            <img
                                src={product.images[selectedImageIndex]}
                                alt={product.name}
                                className="w-full h-full object-cover"
                            />
                        </div>
                    </div>
                    <div className="p-8 md:w-1/2">
                        <h1 className="text-2xl font-bold text-gray-800">
                            {product.name}
                        </h1>
                        <p className="text-gray-600 mt-2">
                            {product.description}
                        </p>

                        <div className="mt-4">
                            <span className="text-gray-700 font-semibold">
                                Price:
                            </span>
                            <span className="text-xl text-green-600 font-bold">
                                {`$${product.price}`}
                            </span>
                        </div>

                        <div className="mt-4">
                            <span className="text-gray-700 font-semibold">
                                Colors:
                            </span>
                            <div className="flex mt-2 space-x-2">
                                {product.colors.map((color, index) => (
                                    <div
                                        key={index}
                                        className={`w-8 h-8 rounded-full border cursor-pointer${
                                            selectedColorIndex === index
                                                ? " ring-2 ring-blue-500"
                                                : ""
                                        }`}
                                        style={{ backgroundColor: color.value }}
                                        title={color.name}
                                        onClick={() => handleColorSelect(index)}
                                    />
                                ))}
                            </div>
                            {product.colors.length > 0 && (
                                <p className="text-sm text-gray-600 mt-1">Selected:<span>{product.colors[selectedColorIndex].name}</span>
                                </p>
                            )}
                        </div>

                        <div className="mt-4">
                            <label className="text-gray-700 font-semibold block mb-2">
                                Quantity:
                            </label>
                            <div className="flex items-center w-max rounded-md overflow-hidden shadow-sm">
                                <button
                                    className="bg-gray-100 hover:bg-gray-200 px-3 py-2 text-gray-700 font-medium cursor-pointer transition-colors"
                                    onClick={decreaseQuantity}
                                >
                                    -
                                </button>
                                <input
                                    type="number"
                                    min="1"
                                    value={quantity}
                                    onChange={(e) =>
                                        setQuantity(
                                            Math.max(
                                                1,
                                                parseInt(e.target.value) || 1,
                                            ),
                                        )
                                    }
                                    className="w-16 text-center py-2 focus:outline-none"
                                    id="quantity-input"
                                />
                                <button
                                    className="bg-gray-100 hover:bg-gray-200 px-3 py-2 text-gray-700 font-medium cursor-pointer transition-colors"
                                    onClick={increaseQuantity}
                                >
                                    +
                                </button>
                            </div>
                        </div>

                        <div className="mt-6">
                            <button
                                className="bg-blue-600 text-white py-2 px-6 rounded-lg font-semibold hover:bg-blue-700 transition duration-200 cursor-pointer"
                                id="add-to-cart-button"
                                onClick={addToCart}
                            >
                                Add to Cart
                            </button>
                        </div>
                    </div>
                </div>

                <div className="p-8">
                    <h2 className="text-xl font-semibold text-gray-800">
                        Product Gallery
                    </h2>
                    <div className="grid grid-cols-3 gap-4 mt-4">{product.images.map((image, index) => (
                            <img
                                key={index}
                                src={image}
                                className={`h-32 w-full object-cover rounded cursor-pointer${
                                    selectedImageIndex === index
                                        ? " ring-2 ring-blue-500"
                                        : ""
                                }`}
                                onClick={() => handleImageSelect(index)}
                            />
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default App;
