# TeamWork

A CRM for your job.

* Remember what's important to your team
* Stay connected and keep your team's needs top-of-mind
* Track achievements for yourself and your team

#### Contents

1. [Quickstart](#quickstart)
2. [Roadmap](#roadmap)

---

## Quickstart

TODO

## Roadmap

### v1

* [ ] CRUD for first class entities (vertex)
    * [x] Create vertex
    * [ ] Retrieve vertex
    * [ ] List vertex
    * [ ] Update vertex
    * [ ] Delete vertex
* [ ] CRUD for edges
    * [x] Create edges
    * [ ] Retrieve edge
    * [ ] List edge
    * [ ] Update edge
    * [ ] Delete edge
* [ ] Storage
    * [ ] Native graph backend
        * [ ] Export to file
        * [ ] Import from file
* [ ] User Interface
    * [ ] HTTP API for CRUD operations
    * [ ] Web UI for CRUD operations

### v1.1

* [ ] Storage
    * [ ] SQLite backend
    * [ ] PostgreSQL backend
    * [ ] Cloud storage integration for DB backups + media
* [ ] User Interface
    * [ ] TUI for CRUD operations
* [ ] Document storage - attached documents to notes as Attr
* [ ] Drop-in replacement for `json`

### v2

* [ ] Plug-in system based on events
* [ ] Users + AuthN/Z as a plugin

### v3

*

---

## Notes

#### 2022.01.30

* # 1 is Plug-ins - a plug-in system that allows others to install logical and UI plugins via config-only would be valuable and neat.
    * Almost everything here could be a plug-in. Including user AuthN/AuthZ.
    * Register on app load - config only, no code changes
    * Plug-ins saved in registry and assigned channels
    * Listen on events over channels, do things at that time
    * Publish things on their own channels, which are listened to at least by teamwork.App
* Multi-user: this is a decently large lift if we want different permission sets, but must be done! The main issue is
  how to segregate Vertex and Edge by User for backends that only support triples - one way is with extra edges from the
  User to the Edges and Vertices.
    * How to handle AuthZ? Definitely Polar as a library.
* Cloud-based storage for local database backup and media (via config-only) would go a long way to making this more
  realistic to use.
* Document storage - if this is a CRM for work, we need to store documents
* Common properties - we should provide convenience methods on our APIs to set common properties, like “position/title”
  or “pronouns” on a Person for a company. Maybe MetaTag, maybe not!
* TODO: use MetaTag for all properties that aren’t explicitly in Attrs?
* Edges - these are roughly quads, but not all graph databases support quads (ex: Dagger). How to handle Label in RDF
  with only triples?
    * We could use an ID-only by default, since we are using ksuid, with support but not requirement for “Type” as
      level.
* Errors - well formed error from all our functions.
    * Cockroach or go-errors
* Code generation - automatically generate code for Typescript types, generated SDK code, …
* GraphQL and REST via HTTP
* UI notification on load for all clients if in-memory database and no backup/save option.
    * Automated backups/save as a first class consideration
* Events - expose /events/:id endpoint that can be queries for some buffer of events (added attr X to entity Y, …)
    * New Event as vertex type
    * ex: Added phone to Person ___
    * We can show a Recent Activity table
