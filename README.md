# Figma to Webpage MCP Server
A Model Context Protocol (MCP) server that converts Figma designs to HTML/CSS webpages using Go.
Features

- ✅ Convert Figma to HTML/CSS - Transform Figma designs into responsive web pages
- ✅ Auto-layout Support - Converts Figma's Auto Layout to CSS Flexbox
- ✅ Style Extraction - Extracts colors, typography, spacing, and borders
- ✅ Component Hierarchy - Maintains design structure in HTML
- ✅ Real-time Access - Fetch designs directly from Figma API
Project Setup
## 1. Prerequisites

- Go 1.21 or higher
- Figma account with API access
- Claude Desktop (for MCP integration)

## 2. Get Your Figma Access Token

- Go to Figma Settings
- Scroll to "Personal access tokens"
- Click "Generate new token"
- Give it a name (e.g., "MCP Server")
- Select scope: file_content:read
- Copy the token (you won't see it again!)