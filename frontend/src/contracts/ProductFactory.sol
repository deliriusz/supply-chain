// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Counters.sol";
import "./ProductLib.sol";
import "./FirmexProductNFT.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

/**
   @title Factory of products
   @author Rafał Kalinowski
   @notice This is just a simple implementation. Please do not use for production purposes without proper audit. Even though I know you won't ;-)
 */
contract ProductFactory is Ownable, IERC721Receiver {
    mapping(uint256 => ProductLib.Product) _products;
    FirmexProductNFT private nft;

    event ProductCreated(
        string indexed _name,
        uint256 _initialPrice,
        uint256 indexed _extId,
        string _nftUri
    );

    event ProductSold(
        uint256 indexed _forPrice,
        uint256 indexed _extId,
        address indexed _to
    );

    modifier productExists(uint256 _id) {
        require(
            _products[_id].initialPrice > 0,
            "Product with given id does not exist"
        );
        _;
    }

    constructor(address nftAddress) {
        nft = FirmexProductNFT(nftAddress);
    }

    function renounceOwnership() public view override onlyOwner {
        revert();
    }

    function transferOwnership(address) public view override onlyOwner {
        revert();
    }

    /**
        @dev IERC721Receiver implementation
     */
    function onERC721Received(
        address,
        address,
        uint256 tokenId,
        bytes calldata
    ) external view override productExists(tokenId) returns (bytes4) {
        return this.onERC721Received.selector;
    }

    function getNftAddress() public view returns (address) {
        return address(nft);
    }

    /**
   @dev Creates an object. Please remember to transfer ownership to this contract for FirmexProductNFT
   @param _name product name
   @param _initialPrice initial price in wei
   @param _extId externalId, used to match product NFT with external DB
 */
    function create(
        string calldata _name,
        uint256 _initialPrice,
        uint256 _extId,
        string calldata _nftUri
    ) external onlyOwner {
        require(
            _products[_extId].initialPrice == 0,
            "Product with given id already exists"
        );
        require(_initialPrice > 0, "Initial Price must be greater than 0");

        ProductLib.Product memory product = ProductLib.Product({
            state: ProductLib.LifecycleState.InProduction,
            initialPrice: _initialPrice,
            name: _name
        });

        _products[_extId] = product;
        nft.safeMint(address(this), _nftUri, _extId);

        emit ProductCreated(_name, _initialPrice, _extId, _nftUri);
    }

    /**
      @dev changes product state, only if product exists and the next state is greater than actual
      @param _extId externalId used to match a product
      @param _newState new lifecycle state
   */
    function changeProductState(
        uint256 _extId,
        ProductLib.LifecycleState _newState
    ) external onlyOwner productExists(_extId) {
        ProductLib.Product storage matchingProduct = _products[_extId];
        require(
            matchingProduct.state < _newState,
            "New state should be bigger than actual"
        );
        matchingProduct.state = _newState;
    }

    /**
      @dev retrieves full product from mapping
      @param _extId externalId used to match a product
   */
    function getProduct(uint256 _extId)
        public
        view
        productExists(_extId)
        returns (ProductLib.Product memory)
    {
        return _products[_extId];
    }

    /**
      @dev If payed amount required for a product matches it price, it transfers NFT to `_to` address, and sends the data to contract owner
      @param _extId externalId used to match a product
      @param _to address to which send product NFT
   */
    function sellProduct(uint256 _extId, address _to)
        public
        payable
        productExists(_extId)
    {
        ProductLib.Product storage product = _products[_extId];

        require(_to != address(0), "destination address cannot be 0");
        require(
            msg.value == product.initialPrice,
            "Value sent for the product does not match product's price"
        );
        require(
            product.state == ProductLib.LifecycleState.Created,
            "Product should be created and not yet payed"
        );
        product.state = ProductLib.LifecycleState.Payed;

        nft.safeTransferFrom(address(this), _to, _extId);
        payable(owner()).transfer(msg.value);

        emit ProductSold(product.initialPrice, _extId, _to);
    }
}
