# Example Use Case: Censorship-Resistant Content Publishing

**Disclaimer:** This document outlines *one theoretical approach* to building a censorship-resistant publishing system using Golem Base. It serves as an illustrative example of how the core functionalities (data storage, retrieval, TTL, annotations) can be leveraged. Actual implementations may vary based on specific requirements and chosen technology stack.

## Scenario

An independent journalist or organization needs a reliable platform to publish articles, reports, or updates, ensuring high availability and resistance to censorship or takedowns often associated with centralized hosting providers. They require a system where content, once published, remains accessible based on its cryptographic identifier, independent of any single server.

## Conceptual Interaction Flow (RPC/Library)

Integrating Golem Base into a publishing application, especially a web-based one, typically involves frontend JavaScript interacting with the user's browser wallet (like MetaMask) and a Golem Base node (e.g., `https://api.golembase.demo.golem-base.io/`) via its JSON-RPC interface. Libraries like Ethers.js or Viem are commonly used to simplify these interactions.

1.  **Wallet Connection & Signing:** The user connects their wallet (e.g., MetaMask) to the application. When an action requires a transaction (like publishing content), the application constructs the transaction parameters and requests the user's wallet to sign and send it. The application *does not* handle the user's private key directly; the wallet manages signing securely.

2.  **Content Preparation:** The article content (text, metadata) needs to be prepared, typically encoded into a byte format (e.g., UTF-8 text converted to hex) by the application before constructing the transaction.

3.  **Store Content via `golembase_createEntity`:** The application constructs and sends a transaction to the Golem Base storage processor address. The transaction's data field contains an RLP-encoded `StorageTransaction` specifying the `Create` operation.

    *   **RPC Method:** `eth_sendRawTransaction` (after signing the transaction locally) or potentially a higher-level library function.
    *   **Underlying Operation:** The transaction targets the Golem Base storage contract, triggering the `golembase_createEntity` logic internally upon mining.

    ```json
    // Conceptual JSON-RPC Request (Illustrative - actual tx is signed and sent)
    // This represents the parameters needed to build the transaction data
    {
      "jsonrpc": "2.0",
      "method": "golembase_createEntity", // Conceptual representation of the target operation
      "params": [{
        "from": "<Publisher's Address>", // Derived from the signing key
        "data": "0x...", // Hex-encoded article content
        "ttl": 2592000, // Example: ~30 days in blocks (adjust as needed)
        "annotations": { // Optional: For indexing and discovery
          "contentType": "article/markdown",
          "publication": "example-journal",
          "authorId": "user-123"
          // Numeric annotations could also be used if needed
        }
      }],
      "id": 1 // Request ID
    }

    // Upon successful transaction mining, the application would typically monitor
    // for the GolemBaseStorageEntityCreated event log to retrieve the generated entityKey.
    // Event Log Topic[1] contains the entityKey (bytes32).
    ```
    *   **Result:** The network assigns a unique `entityKey` (a `bytes32` hash) to the stored content. The application needs to capture this key, often by parsing transaction receipts and logs.

4.  **Content Distribution:** The publisher shares the `entityKey` (e.g., `0xabc...def`). This key acts as the permanent, decentralized identifier for the article.

5.  **Retrieve Content via `golembase_getEntity`:** Readers or client applications use the `entityKey` to fetch the content directly from any Golem Base node.

    ```json
    // Conceptual JSON-RPC Request
    {
      "jsonrpc": "2.0",
      "method": "golembase_getEntity", // Direct RPC call to fetch data
      "params": ["<entityKey>"], // The unique key obtained in step 3
      "id": 2
    }
    // Expected Response:
    {
      "jsonrpc": "2.0",
      "result": {
        "key": "<entityKey>",
        "owner": "<Publisher's Address>",
        "expiration": "<block number>",
        "payload": "0x...", // Hex-encoded article content
        "stringAnnotations": [ // Annotations associated with the entity
           {"key": "contentType", "value": "article/markdown"},
           {"key": "publication", "value": "example-journal"},
           {"key": "authorId", "value": "user-123"}
        ],
        "numericAnnotations": []
      },
      "id": 2
    }
    ```
    *   The client application receives the hex-encoded payload and decodes it for display.*

---

**Note on the `golembase-demo-cli`:**

The command-line tool (`golembase-demo-cli`) simplifies this workflow for demonstration and testing. Commands like `golembase-demo-cli entity create --data "..." --ttl ... --annotations '{"key": "value"}'` abstract the complexities of:
*   Loading the user's private key.
*   Encoding the storage operation (Create) into RLP format.
*   Constructing the Ethereum transaction (including nonce, gas estimation, chain ID).
*   Signing the transaction.
*   Sending it via `eth_sendRawTransaction`.
*   Waiting for the transaction receipt and extracting the `entityKey` from logs.

Similarly, `golembase-demo-cli cat <entityKey>` handles the `golembase_getEntity` RPC call and decodes the result. While extremely useful for demos, production systems require direct integration using RPC or dedicated libraries for robust error handling, key management, and user interface development.
