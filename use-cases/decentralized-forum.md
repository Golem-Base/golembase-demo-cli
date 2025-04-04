# Example Use Case: Decentralized Forum / Bulletin Board System

**Disclaimer:** This document illustrates *one theoretical approach* to building a decentralized forum or bulletin board system (BBS) using Golem Base primitives. It demonstrates how features like annotated data storage and querying can be combined. Actual implementations will likely require more complex logic for threading, user profiles, moderation, etc., and may vary based on the chosen architecture.

## Scenario

A community aims to establish an online forum where members can create discussion topics, post replies, and discover relevant conversations using categories or tags. The goal is to create a platform resistant to single points of failure or censorship, where the discussion data persists on the Golem Base network.

## Conceptual Interaction Flow (RPC/Library)

Building such a forum involves storing posts (topics and replies) as distinct entities and leveraging annotations for organization, discovery, and potentially linking replies to topics. Interactions, especially in a web application, would typically use frontend JavaScript with libraries like Ethers.js or Viem to interact with the user's browser wallet (e.g., MetaMask) and a Golem Base node (e.g., `https://api.golembase.demo.golem-base.io/`).

1.  **Wallet Connection & Signing:** Users connect their wallet to the forum application. When posting a topic or reply, the application prepares the necessary transaction details (operation type, data, annotations, TTL) and prompts the user's wallet to sign and broadcast the transaction. The application itself does not handle private keys.

2.  **Create a Forum Topic:** A user initiates a new discussion by creating an entity representing the first post of the topic. Annotations are crucial for categorization and discovery.

    *   **RPC Method:** `eth_sendRawTransaction` (after signing the transaction locally).
    *   **Underlying Operation:** Targets the Golem Base storage contract, triggering `golembase_createEntity`.

    ```json
    // Conceptual JSON-RPC Request (Illustrative - parameters for tx data)
    {
      "jsonrpc": "2.0",
      "method": "golembase_createEntity", // Conceptual target operation
      "params": [{
        "from": "<User's Address>",
        "data": "0x...", // Hex-encoded post content: "Discussion: Best use cases for L3 DB Chains?"
        "ttl": 5184000, // Example: ~60 days in blocks
        "annotations": { // Structured annotations for querying
          "app": "forum-v1", // Identify the application
          "type": "topic", // Distinguish topics from replies
          "category": "technical-discussion",
          "tags": "L3,data,scalability" // Use consistent tagging
          // "topicId": "<generated-unique-topic-id>" // Could be the entityKey itself or another identifier
        }
      }],
      "id": 1
    }
    // Upon success, monitor logs for GolemBaseStorageEntityCreated to get the entityKey.
    // This entityKey can serve as the unique ID for this topic.
    ```
    *   **Result:** A new entity is created, representing the topic's initial post. Its `entityKey` is captured by the application.

3.  **Discover Topics via `golembase_queryEntities`:** Other users can find topics using queries based on annotations.

    ```json
    // Conceptual JSON-RPC Request to find topics in 'technical-discussion'
    {
      "jsonrpc": "2.0",
      "method": "golembase_queryEntities",
      "params": ["app = 'forum-v1' AND type = 'topic' AND category = 'technical-discussion'"], // Query string
      "id": 2
    }
    // Expected Response:
    {
      "jsonrpc": "2.0",
      "result": [ // Array of SearchResult objects
        { "key": "<key1>", "value": "0x...", "stringAnnotations": [...], "numericAnnotations": [...] },
        { "key": "<key2>", "value": "0x...", "stringAnnotations": [...], "numericAnnotations": [...] }
        // Potentially includes payload and annotations directly in results
      ],
      "id": 2
    }
    ```
    *   The application parses the results to display a list of matching topics.*

4.  **Retrieve Topic/Post Content via `golembase_getEntity`:** When a user selects a topic, the application fetches the full entity data (including the post content) using its `entityKey`.

    ```json
    // Conceptual JSON-RPC Request
    {
      "jsonrpc": "2.0",
      "method": "golembase_getEntity",
      "params": ["<entityKey>"], // Key of the topic's initial post
      "id": 3
    }
    // Expected Response: (Similar structure as in the publishing use case, containing payload, owner, annotations etc.)
    ```

5.  **(Optional) Posting Replies:** Replies could be implemented as separate entities linked to the original topic via an annotation (e.g., `replyToTopic: "<topicEntityKey>"`). Discovering replies for a topic would involve querying for entities with `type = 'reply'` and the matching `replyToTopic` annotation.

---

**Note on the `golembase-demo-cli`:**

The command-line tool (`golembase-demo-cli`) provides commands that map to these underlying operations, simplifying testing:
*   `golembase-demo-cli entity create --data "..." --annotations '{"app": "forum-v1", ...}'`: Creates a topic or reply entity.
*   `golembase-demo-cli query "type = 'topic' AND category = 'technical-discussion'"`: Executes annotation-based queries.
*   `golembase-demo-cli cat <entityKey>`: Retrieves the content of a specific post.

While the CLI demonstrates the core capabilities, building a functional forum requires a dedicated application to handle user interfaces, manage relationships between topics and replies, implement user profiles, and potentially add moderation features, all interacting with the Golem Base RPC interface or libraries.
