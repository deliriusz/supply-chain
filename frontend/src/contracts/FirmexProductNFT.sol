// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/// @custom:security-contact kalinowski.software@gmail.com
contract FirmexProductNFT is ERC721, ERC721URIStorage, Ownable {
    string internal baseUrl;

    constructor(string memory _baseURL) ERC721("FirmexProductNFT", "FXP") {
        baseUrl = _baseURL;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseUrl;
    }

    function safeMint(
        address to,
        string memory uri,
        uint256 id
    ) public onlyOwner {
        uint256 tokenId = id;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }

    // The following functions are overrides required by Solidity.

    function _burn(uint256 tokenId)
        internal
        override(ERC721, ERC721URIStorage)
    {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }
}
