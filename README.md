# **Loci Proto** – gRPC Protocol Definitions 🔄✨

This repository contains the **Protocol Buffer (protobuf)** definitions for Loci's microservices architecture. These proto files define the gRPC services and message types that power Loci's AI-driven personalized city discovery platform.

## 🚀 About Loci

**Loci** is a smart, mobile-first web application delivering hyper-personalized city exploration recommendations based on user interests, time, location, and an evolving AI engine. It combines advanced AI personalization with real-time streaming capabilities to provide contextual, intelligent city discovery experiences.

### Core Value Proposition
Tired of generic city guides? Loci learns your preferences (history, food, art, etc.) and combines them with your available time and location to suggest the perfect spots, activities, and restaurants.

## 📁 Repository Structure

```
go-ai-poi-proto/
├── proto/                    # Protocol Buffer definitions
│   ├── auth.proto           # Authentication & authorization
│   ├── chat.proto           # AI-powered chat with streaming
│   ├── poi.proto            # Points of interest discovery
│   ├── user.proto           # User profiles & preferences
│   ├── list.proto           # Lists & itinerary management
│   ├── city.proto           # City information & statistics
│   ├── statistics.proto     # Analytics with real-time streaming
│   ├── recents.proto        # Recent user interactions
│   ├── review.proto         # Reviews & ratings system
│   ├── common.proto         # Shared types & utilities
│   └── ai_poi_service.proto # Main API gateway service
├── modules/                 # Generated gRPC modules
├── container/               # Dependency injection
├── core/                    # Core gRPC infrastructure
└── utils/                   # Connection & transport utilities
```

## 🌟 Key Features

### 🔄 **Streaming Services**
- **AI Chat Streaming**: Real-time conversation with Gemini-powered AI assistant
- **Live Statistics**: Real-time metrics and analytics updates
- **Progressive Results**: Streaming discovery results for better UX

### 🧠 **AI-Powered Services**
- **Semantic Search**: Vector embeddings with PostgreSQL `pgvector`
- **Personalized Recommendations**: Based on user preferences and behavior
- **Contextual Filtering**: Time, location, interests, and availability-aware
- **Smart Itinerary Planning**: AI-generated travel plans

### 🗺️ **Core Discovery Services**
- **Multi-Domain Search**: POIs, restaurants, hotels, attractions
- **Geospatial Queries**: PostGIS-powered location services
- **Hybrid Search**: Combines spatial and semantic search
- **Real-time Filtering**: Opening hours, distance, rating, price

## 🔧 Service Architecture

```mermaid
graph TB
    Gateway[AI POI Gateway] --> Auth[AuthService]
    Gateway --> Chat[ChatService - STREAMING]
    Gateway --> POI[POIService]
    Gateway --> User[UserService]
    Gateway --> List[ListService]
    Gateway --> City[CityService]
    Gateway --> Stats[StatisticsService - STREAMING]
    Gateway --> Recents[RecentsService]
    Gateway --> Review[ReviewService]
    
    Chat -.->|Server-Sent Events| Client[Web/Mobile Clients]
    Stats -.->|Real-time Updates| Client
    
    subgroup Backend Services
        Auth --> DB[(PostgreSQL + PostGIS)]
        POI --> DB
        Chat --> Gemini[Google Gemini API]
        POI --> Vector[pgvector Embeddings]
    end
```

## 🛠️ Technology Integration

### Backend Stack Compatibility
- **Go Backend**: `go-ai-poi-server` with Chi/Gin Gonic routers
- **Database**: PostgreSQL with PostGIS for geospatial queries
- **AI Engine**: Google Gemini API integration via `google/generative-ai-go`
- **Vector Search**: PostgreSQL with `pgvector` extension
- **Authentication**: JWT tokens with `Goth` package for social logins

### Frontend Integration
- **Web**: SvelteKit client (`go-ai-poi-client`)
- **Mobile**: iOS app (`go-ai-poi-ios`) and Angular PWA (`go-ai-poi-ng`)
- **Real-time**: Server-Sent Events (SSE) for streaming responses

## 📡 Streaming Implementation

### Chat Service Streaming
```protobuf
service ChatService {
  // Real-time AI conversations
  rpc StartChatStream(StartChatRequest) returns (stream ChatEvent);
  rpc ContinueChatStream(ContinueChatRequest) returns (stream ChatEvent);
  rpc FreeChatStream(FreeChatRequest) returns (stream ChatEvent);
}
```

### Statistics Service Streaming
```protobuf
service StatisticsService {
  // Live metrics updates
  rpc StreamMainPageStatistics(StreamMainPageStatisticsRequest) returns (stream StatisticsEvent);
}
```

## 🔐 Authentication & Security

- **JWT-based Authentication**: Secure token-based auth
- **Social Login Support**: Google, Facebook, Apple integration
- **Role-based Access**: User, premium, admin permissions
- **Rate Limiting**: Built-in request throttling
- **Input Validation**: Comprehensive request validation

## 💰 Business Model Integration

### Freemium Features
- **Free Tier**: Core recommendations, basic filters, limited saves
- **Premium Tier**: Enhanced AI, advanced filters, unlimited saves, offline access

### Monetization Endpoints
- **Partnership APIs**: Booking referrals, commission tracking
- **Featured Listings**: Premium business visibility
- **Analytics APIs**: Anonymized trend data for tourism boards

## 🚀 Getting Started

### Prerequisites
```bash
# Install Protocol Buffer Compiler
brew install protobuf  # macOS
# or
apt-get install protobuf-compiler  # Ubuntu

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Code Generation
```bash
# Generate Go code from proto files
make generate

# Or manually:
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/*.proto
```

### Development Setup
```bash
# Clone the repository
git clone https://github.com/your-org/go-ai-poi-proto.git
cd go-ai-poi-proto

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build modules
make build
```

## 📊 API Coverage

| Service | Endpoints | Streaming | Description |
|---------|-----------|-----------|-------------|
| **AuthService** | 8 | ❌ | Authentication, OAuth, JWT management |
| **ChatService** | 7 | ✅ | AI chat, itinerary planning, streaming |
| **POIService** | 15 | ❌ | Discovery, search, recommendations |
| **UserService** | 20 | ❌ | Profiles, preferences, interests, tags |
| **ListService** | 15 | ❌ | Lists, itineraries, favorites |
| **CityService** | 4 | ❌ | City data, statistics, search |
| **StatisticsService** | 6 | ✅ | Analytics, metrics, real-time updates |
| **RecentsService** | 5 | ❌ | User activity, interaction history |
| **ReviewService** | 10 | ❌ | Reviews, ratings, business responses |

**Total**: 90+ gRPC endpoints covering the entire Loci platform

## 🌍 Multi-Client Support

### Web Client (`go-ai-poi-client`)
- SvelteKit with TypeScript
- Real-time chat interface
- Progressive discovery results
- Mobile-responsive design

### iOS App (`go-ai-poi-ios`)
- Swift with SwiftUI
- Native gRPC integration
- Offline-first architecture
- Push notifications

## 🔄 Migration from REST

This protobuf repository supports the migration from REST to gRPC:

### Phase 1: Hybrid Architecture
- REST endpoints remain active
- gRPC services run in parallel
- Gradual client migration

### Phase 2: Streaming Features
- Chat streaming via gRPC
- Real-time statistics
- Progressive loading

### Phase 3: Full Migration
- Complete REST deprecation
- Pure gRPC architecture
- Enhanced performance

## 🧪 Development & Testing

### Proto Validation
```bash
# Validate proto files
buf lint proto/
buf breaking proto/ --against .git#branch=main
```

### Service Testing
```bash
# Test gRPC services
grpcurl -plaintext localhost:9090 list
grpcurl -plaintext -d '{"message": "Hello"}' localhost:9090 ai_poi.chat.v1.ChatService/FreeChatStream
```

### Load Testing
```bash
# Performance testing
ghz --insecure \
    --proto proto/chat.proto \
    --call ai_poi.chat.v1.ChatService.StartChatStream \
    --data '{"user_id":"test","initial_message":"Hello"}' \
    localhost:9090
```

## 📋 Roadmap

### Phase 1 (Current)
- [x] Core service definitions
- [x] Authentication & user management
- [x] Basic POI discovery
- [x] Chat streaming implementation

### Phase 2 (In Progress)
- [ ] Advanced AI features (embeddings, semantic search)
- [ ] Multi-language support
- [ ] Enhanced analytics
- [ ] Business partnership APIs

### Phase 3 (Planned)
- [ ] Voice interface integration
- [ ] Augmented reality features
- [ ] Multi-city expansion
- [ ] Real-time collaborative planning

## 🤝 Contributing

We welcome contributions to improve Loci's gRPC definitions!

### Guidelines
1. **Follow protobuf best practices**
2. **Maintain backward compatibility**
3. **Add comprehensive documentation**
4. **Include validation rules**
5. **Write tests for new services**

### Pull Request Process
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and validation
5. Submit a pull request

## 📄 License

This project is part of the Loci ecosystem. License details coming soon.

---

## 🔗 Related Repositories

- **[go-ai-poi-server](../go-ai-poi-server)** - Go backend services
- **[go-ai-poi-client](../go-ai-poi-client)** - SolidStart web application  
- **[go-ai-poi-ios](../go-ai-poi-ios)** - iOS native application
- **[go-ai-genai-sdk](../go-ai-genai-sdk)** - AI/ML SDK integration

---

*Built with ❤️ for intelligent city discovery*
