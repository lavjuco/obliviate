## Introduction

This README provides step-by-step instructions on how to set up the preset.json configuration file. Follow the steps below to obtain the required tokens, author/user ID, and channel/target IDs.

### `preset.json` Example:
```json
{
  "token": "user_token",
  "author": "user_id",
  "channels": ["target_id_1", "target_id_2"]
}
```

## How to Get Author/User ID:

1. Open Discord in your browser.
2. Enable Developer Mode:
    + Go to User Settings.
    + Under the "Advanced" tab, enable "Developer Mode."
3. Right-click on yourself.
4. Select "Copy User ID."

## How to Get Channel/Target ID:

1. Open Discord in your browser.
2. Enable Developer Mode (follow step 2 above).
3. For DM (Direct Message):
    + Open the DM.
    + Right-click on the user's name.
    + Select "Copy Channel ID."
4. For Server:
    + Navigate to the server.
    + Right-click on the server icon.
    + Select "Copy Server ID."

## How to Get Token:

1. Open Discord in your browser.
2. Open Web Developer Tools:
    + Right-click anywhere on the page.
    + Select "Inspect" or "Inspect Element."
    + Go to the "Network" tab.
3. Reload the Discord page to capture network requests.
4. Select a request (e.g., https://discord.com/api/v9/science).
5. In the Headers section, find and copy the Authorization token.

Ensure that you keep your token confidential and do not share it with others. Use the obtained information to fill in the corresponding fields in the preset.json file for successful setup.
