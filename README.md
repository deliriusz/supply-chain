# supply-chain
This Work-In-Progress project is just a exemplary implementation of an online shop and a simple blockchain based supply chain for it created as a proof of concept.

Here are main features of this project:
* onchain payments
* blockchain based user log in and admin RBAC
* minting products as NFTs
* maintaining and editing product catalog in DB
* statistics chart for admins
* onchain supply chain tracking

## Setup

### Frontend
```
cd {root}/frontend
npm install
npx hardhat compile
npm start
```

### Backend
```
cd {root}/backend
go install
make eth
make run
```

You also need to set proper variables in ***.env*** files for both backend and frontend
## Testing
### Frontend
```
cd {root}/frontend
npm test
```
### Backend
```
cd {root}/backend
make test
```