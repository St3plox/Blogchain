

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Blogchain </h3>

  <p align="center">
    A simple blockchain-based blogging platform
    <br />
<!--     <a href="https://github.com/St3plox/Blogchain"><strong>Explore the docs »</strong></a> -->
    <br />
    <br />
    <a href="https://github.com/St3plox/Blogchain/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/St3plox/Blogchain/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
    <a href="https://pkg.go.dev/github.com/St3plox/Blogchain">Documentation</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project was created mostly for learning purposes. The app itself represents a frontend and backend api that uses document database and ethereum testnet. 
When user creates an account the ethereum address will be generated and asociated with their account. All user data are stored in the db. However, posts
are stored in ethereum blockchain. There is also an admin account that signs all the transactions and pays the gas fees. Backend api uses jwt to authenticate users. 

You can access api docs on http://localhost:4000/swagger/index.html

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Golang][Golang]][Golang-url]
* [![Vue.js][Vue.js]][Vue-url]
* [![JavaScript][JavaScript]][JavaScript-url]
* [![Solidity][Solidity]][Solidity-url]
* [![MongoDB][MongoDB]][MongoDB-url]
* [![Redis][Redis]][Redis-url]
* [![Hardhat][Hardhat]][Hardhat-url]
* [![Abigen][Abigen]][Abigen-url]
* [![Ethereum][Ethereum]][Ethereum-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started
The simplest way to get started is to use docker. Otherwise you will have to install Go 1.22, npm, hardhat, mongod, redis. Startup the mongodb server, than hardhat tesnet, 
than frontend and backend api

### Prerequisites
*docker
*docker-compose
*openssl

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/St3plox/Blogchain.git
   cd Blogchain
   ```
3. Generate private key that is used in auth
   ```bash
   make gen-private
   ```
4. build and run the app
   ```bash
   docker-compose up -d
   ```
Frontend can be accessed on port 8080, backend - 5000

<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- ROADMAP -->
## Roadmap

- [X] Caching
- [X] Documentation
- [ ] Media support
- [ ] Adequate testing
- [ ] CI/CD (or smth like this)
- [ ] Recomendation engine
- [ ] Improved UI/UX (no way)

See the [open issues](https://github.com/St3plox/Blogchain/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Egor - st3pegor@gmail.com


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Ardan labs Github](https://github.com/ardanlabs)
* [Boostrap Vue navbar](https://bootstrap-vue.org/docs/components/navbar)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://golang.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[JavaScript]: https://img.shields.io/badge/javascript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black
[JavaScript-url]: https://www.javascript.com/
[Solidity]: https://img.shields.io/badge/solidity-363636?style=for-the-badge&logo=solidity&logoColor=white
[Solidity-url]: https://docs.soliditylang.org/
[MongoDB]: https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white
[MongoDB-url]: https://www.mongodb.com/
[Hardhat]: https://img.shields.io/badge/Hardhat-FFCF24?style=for-the-badge&logo=hardhat&logoColor=black
[Hardhat-url]: https://hardhat.org/
[Abigen]: https://img.shields.io/badge/abigen-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Abigen-url]: https://pkg.go.dev/github.com/ethereum/go-ethereum/accounts/abi/bind
[Ethereum]: https://img.shields.io/badge/ethereum-3C3C3D?style=for-the-badge&logo=ethereum&logoColor=white
[Ethereum-url]: https://ethereum.org/
[Redis]: https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white
[Redis-url]: https://redis.io/
