---
title: "Cosmos-sdk docs suck"
date: 2022-05-07T11:32:18+02:00
draft: false
---

Cosmos-sdk is a great framework for building blockchains for the cosmos ecosystem, but the docs just suck.

## Cosmos-sdk
First of all there ARE rather in depth [docs](https://docs.cosmos.network) for cosmos, but the flaw with those docs is that they describe the inner workings of a certain feature in depth, but don't show how to actually use them.

Take the [gov](https://docs.cosmos.network/main/modules/gov) module as an example, in cosmos governance can be used to let the userbase of the chain decide on the future of the chain, for example by changing parameters or adding new code snippets. So it might be very obvious, that programmers might want to introduce custom governance votes for their blockchain, so that they can let the users decide stuff that's not covered by the default proposals already present.

This sounds easier then it is and isn't at all documented, the docs just show how governance and the already present proposals work, and that's also very interesting to know, but does not help developing on ones own project.

So to actually find out how those things are done reading the actual sourcecode is required, because an API documentation is not present either.

This does not just apply to gov but to most other features and modules.

## Finding the docs
You might think "a quick google will find me my docs", but that's wrong, the actual [docs](https://docs.cosmos.network) are not even findable via duck duck go and can only be find by navigating to them via the cosmos-sdk homepage. And also googling `Cosmos-sdk proplem XY` does not find anything related to cosmos-sdk but the first then results are related to Cosmos DB from Microsoft.

## Ignite (former Starport)
Ignite is a tool that's most often used for building chains with Cosmos and does most of the development work, so that the programmer in most cases just has to add logic to scaffolded feature.

Ignites docs are not as bad in it self, but finding them or searching anything related to it definitely is. Back when Ignite was still called Starport most google searched returned stuff related to video games and to the tool itself. You might think after renaming Starport to Ignite this might have gotten better, but no. One might think the new name might tend to be distinguishable from other tech related topics, but no. There are several other tools called Ignite related to programming and specially to JS.
So picking this name, not only made googling for Cosmos-sdk, but also for JS programmers worse.

TL;DR Pick names wisely!

