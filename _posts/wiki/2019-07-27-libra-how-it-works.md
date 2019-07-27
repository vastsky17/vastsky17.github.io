---
layout: post
title: Facebook的Libra “区块链”到底是如何运作的？
category: Wiki
tags: 区块链 Libra
keywords: Facebook的Libra,区块链
description: Facebook的Libra “区块链”到底是如何运作的？
coverage: libra_coverage.png
---


本文深入研究了“关于Facebook Libra coin (以及更多)平台协议”的26页技术文档，并对其内容进行了分解说明.同时，我们对这53位作者表示衷心的钦佩！


以下为具体分析内容：


（文中英文内容为“协议”原文，中文翻译是对“协议”内容的解读.）


---


## 摘要

> The Libra protocol allows a set of replicas—referred to as
> 
> 
> validators—from different authorities to jointly maintain a database
> 
> 
> of programmable resources.

换句话说，也就是这个系统需要由一组权威机构以自上而下的方式进行控制.然而，请注意，该数据库是为维护“可编程资源”而不仅仅是维护数字货币的.

> These resources are owned by different user accounts authenticated by
> 
> 
> public key cryptography and adhere to custom rules specified by the
> 
> 
> developers of these resources.

使用诸如“资源”（resources）之类的通用词汇使我怀疑这里不仅仅是指一种稳定币.

> Transactions are based on predefined and, in future versions,
> 
> 
> user-defined smart contracts in a new programming language called
> 
> 
> Move. We use Move to define the core mechanisms of the blockchain,
> 
> 
> such as the currency and validator membership.

好了，这个有意思了.使用专门的智能契约语言会导致很多问题，比如该语言的功能丰富度，以及延伸到该系统对对抗性契约的健壮性有多强的问题.还有一些关于开发人员友好性以及Libra如何保护智能合约开发人员不受影响的问题都是需要明晰的.

> These core mechanisms enable the creation of a unique governance
> 
> 
> mechanism that builds on the stability and reputation of existing
> 
> 
> institutions in the early days but transitions to a fully open system
> 
> 
> over time.

关于开发人员友好性以及Libra如何保护智能合约开发人员不受影响，这仍是问题.


## Facebook Libra

> This ecosystem will offer a new global currency—the Libra coin—which
> 
> 
> will be fully backed with a basket of bank deposits and treasuries
> 
> 
> from high-quality central banks.

Libra是一种通用的加密资产协议，第一个资产将是一种稳定币.

> Over time, membership eligibility will shift to become completely open
> 
> 
> and based only on the member’s holdings of Libra.

听起来很像股权证明.显然，计划是在五年后开放会员资格，并希望他们当时能够找到股份证明——尽管我预计它们会遇到与Ethereum相同的问题.

> The association has published reports outlining … the roadmap for the
> 
> 
> shift toward a permissionless system.

我很确定这将是分布式网络首次从许可型转换为非许可型.也许整个网络可以转换为股权证明，但为了稳定币/篮子，一些实体必须保持对传统金融系统的开放.这将是通过Libra协会长期集中控制的重点.

> Validators take turns driving the process of accepting transactions.
> 
> 
> When a validator acts as a leader, it proposes transactions, both
> 
> 
> those directly submitted to it by clients and those indirectly
> 
> 
> submitted through other validators, to the other validators. All
> 
> 
> validators execute the transactions and form an authenticated data
> 
> 
> structure that contains the new ledger history. The validators vote on
> 
> 
> the authenticator for this data structure as part of the consensus
> 
> 
> protocol.

这听起来像Practical Byzantine Fault Tolerance(实用拜占庭容错算法)，这是一个很好理解的发展了20年的算法，尽管他们可能做了一些调整.我们在白皮书的第5节中了解到它被称为LibraBFT，它是HotStuff共识协议的变体.

> As part of committing a transaction T i at version i, the consensus
> 
> 
> protocol outputs a signature on the full state of the database at
> 
> 
> version i—including its entire history—to authenticate responses to
> 
> 
> queries from clients.

这是值得注意的，主要是因为它意味着新的验证者应该能够加入网络并快速同步，而不必回溯区块链的整个历史记录，前提是它们信任现有的验证者.


这种帐户模型是有可能的，因为Facebook不太可能关注隐私，而它确实对智能合约感兴趣.


## 逻辑数据模型 Model

> The Libra protocol uses an account-based data model to encode the
> 
> 
> ledger state.

从数据结构的角度来看，Libra更像Ethereum或Ripple，而不是比特币.UTXO模型有优点也有缺点——由于基于输出的历史记录的简单性，它具有更好的私密性和更健壮的事务历史记录——但是处理复杂的智能合约可能更困难.因此，账户模式是有意义的，因为Facebook不太可能关注隐私，尽管听起来它对智能合同很感兴趣.

> The Libra protocol does not link accounts to a real-world identity. A
> 
> 
> user is free to create multiple accounts by generating multiple
> 
> 
> key-pairs. Accounts controlled by the same user have no inherent link
> 
> 
> to each other. This scheme follows the example of Bitcoin and Ethereum
> 
> 
> in that it provides pseudonymity for users.

这听起来好得惊人，但我想知道Libra coin是否也是这种情况.对于那些想要开发一些更能保护隐私的应用程序的开发人员来说，观察这个系统的开放程度将是一件很有趣的事情.

> Every resource has a type declared by a module. Resource types are
> 
> 
> nominal types that consist of the name of the type and the name and
> 
> 
> address of the resource’s declaring module.

看起来你可以生成一个地址，只要每个资产都有唯一的名称，该地址就可以分配任意数量的资产.

> Executing a transaction T i produces a new ledger state S i as well as
> 
> 
> the execution status code, gas usage, and event list.

好了，现在我们知道了如何保护系统免受资源耗尽攻击，大概是利用类似于Ethereum的资源成本系统.

> There is no concept of a block of transactions in the ledger history.

有趣.Libra协议中没有实际的区块链数据结构——块更像是一个虚拟的逻辑结构，验证者使用它来协调系统状态的确认快照.回过头来看，这一节的第一句话现在有了更多的意义:

> All data in the Libra Blockchain is stored in a single versioned
> 
> 
> database. A version number is an unsigned 64-bit integer that
> 
> 
> corresponds to the number of transactions the system has executed.

我所熟悉的每个加密资产网络都以相同的方式在非常高的层次上工作：首先存在一个系统状态，然后执行一个事务，实际上是一个状态转换函数，接着新的系统状态就出现了.



将批量事务放入容器或块中的目的是为了对它们进行排序和加时间戳.这对于无许可网络非常重要，在这种网络中，数据通过动态多方成员签名进行身份验证，验证者可以自由地加入和离开网络.因为Libra运行一个经过许可的系统，所以它可以使用一个更有效的协商一致算法，而不需要批处理事务，因为事务历史记录被重写的可能性要小得多.

> In the initial version of the Libra protocol, only a limited subset of
> 
> 
> Move’s functionality is available to users. While Move is used to
> 
> 
> define core system concepts, such as the Libra currency, users are
> 
> 
> unable to publish custom modules that declare their own resource
> 
> 
> types. This approach allows the Move language and toolchain to
> 
> 
> mature—informed by the experience in implementing the core system
> 
> 
> components—before being exposed to users. The approach also defers
> 
> 
> scalability challenges in transaction execution and data storage that
> 
> 
> are inherent to a general-purpose smart contract platform.

这听起来非常类似于前面提到的“open validator membership（开放验证者成员资格）”计划.似乎Facebook还没有解决任何一个Ethereum多年来一直在努力解决的重大问题.

> In order to manage demand for compute capacity, the Libra protocol
> 
> 
> charges transaction fees, denominated in Libra coins.

Libra coins实际上是协议的原生单位，就像ETH是Ethereum的原生单位.这就引出了另一个关于Libra匿名性质的问题：你可以在没有AML / KYC的情况下获得币吗？如果不能，那么您似乎无法匿名地使用系统的任何功能.查阅Calibra钱包，它将需要AML / KYC.所以我想知道最终是否会有一些进入系统的方式没有受到严格控制.

> The system is designed to have low fees during normal operation, when
> 
> 
> sufficient capacity is available.

这确实很模糊，并引发了许多问题：什么是低收费？什么是正常操作？什么是足够的容量？


## 执行交易 Move

> Many parts of the core logic of the blockchain are defined using Move,
> 
> 
> including the deduction of gas fees. To avoid circularity, the VM
> 
> 
> disables the metering of gas during the execution of these core
> 
> 
> components.

这听起来很危险，但该文档的作者指出，核心组件必须以防御性方式编写以防止DoS攻击.

> The key feature of Move is the ability to define custom resource types
> 
> 
> … the Move type system provides special safety guarantees for
> 
> 
> resources. A resource can never be copied, only moved. These
> 
> 
> guarantees are enforced statically by the Move VM. This allows us to
> 
> 
> represent Libra coins as a resource type in the Move language.

这就澄清了之前的问题：Libra coins是否像ETH或BTC一样是本地资产.我希望这些币只是系统启动时默认的或唯一允许的资源类型，其他资源将在未来提供.

> Move’s stack-based bytecode has fewer instructions than a higher-level
> 
> 
> source language would. In addition, each instruction has simple
> 
> 
> semantics that can be expressed via an even smaller number of atomic
> 
> 
> steps. This reduces the specification footprint of the Libra protocol
> 
> 
> and makes it easier to spot implementation mistakes.

这听起来像是经过深思熟虑的; 希望这意味着他们的脚本语言的安全性将比Ethereum更好.


我们看到“Libra区块链” 实际上并不是区块链.


## 已验证的数据结构和存储 Merkle

> The Libra protocol uses a single Merkle tree to provide an
> 
> 
> authenticated data structure for the ledger history … specifically,
> 
> 
> the ledger history uses the Merkle tree accumulator approach to form
> 
> 
> Merkle trees, which also provides efficient append operations.

我们再一次看到“Libra区块链”实际上并不是区块链.这个协议似乎设计得非常好，但是奇怪的是，当账户历史的数据结构是一组有签名的账户状态时，它们仍然称它为区块链.验证者正在为每个账户状态做出承诺，并且所有历史帐户状态也都在Merkle树中承诺，但我还没有真正看到形成链的任何反向链接数据列表——更不用说形成块链了.

> The authenticator of an account is the hash of this serialized
> 
> 
> representation. Note that this representation requires recomputing the
> 
> 
> authenticator over the full account after any modification to the
> 
> 
> account. The cost of this operation is O(n), where n is the length of
> 
> 
> the byte representation of the full account.

嗯，如果没有对给定帐户存储的数据量进行限制，这听起来像是DoS攻击的开端.

> We anticipate that as the system is used, eventually storage growth
> 
> 
> associated with accounts may become a problem. Just as gas encourages
> 
> 
> responsible use of computation resources, we expect that a similar
> 
> 
> rent-based mechanism may be needed for storage. We are assessing a
> 
> 
> wide range of approaches for a rent-based mechanism that best suits
> 
> 
> the ecosystem.

另一个未解决的问题.迫不及待地想说“租金太高了！”

> The voting power must remain honest both during the epoch as well as
> 
> 
> for a period of time after the epoch in order to allow clients to
> 
> 
> synchronize to the new configuration. A client that is offline for
> 
> 
> longer than this period needs to resynchronize using some external
> 
> 
> source of truth to acquire a checkpoint that they trust.

哎.目前尚不清楚这个“时间段”有多长，但如果一个epoch不到一天，那么我猜测指定的“时间段”也是如此.看起来这个共识协议不够强大，参与者可能会随意离开并重新加入网络.


## 拜占庭容错共识 Byzantine

> LibraBFT assumes that a set of 3f + 1 votes is distributed among a set
> 
> 
> of validators that may be honest, or Byzantine. LibraBFT remains safe,
> 
> 
> preventing attacks such as double spends and forks when at most f
> 
> 
> votes are controlled by Byzantine validators.

就像PBFT一样，这种一致性算法可以容忍33％的验证者是不诚实的.HotStuff的修改听起来很合理：


通过使验证者签署块的状态（而不仅仅是事务序列）来抵制非确定性错误.


一个发出明确超时信号的起搏器，验证者依赖于这些超时信号的仲裁集来进入下一轮 - 这应该可以提高活性.


不可预知的领导者选举机制，以限制针对领导者的DoS攻击.


聚合签名以便保存那些签署了仲裁集证书来为块接受投票的身份验证者.


## 网络 NetWork

> Each validator in the Libra protocol maintains a full membership view
> 
> 
> of the system and connects directly to any validator it needs to
> 
> 
> communicate with. A validator that cannot be connected to directly is
> 
> 
> assumed to fall within the quota of Byzantine faults tolerated by the
> 
> 
> system.

这将需要大量工作才能将系统扩展到数百个验证者.

1. Libra核心实施内容

---

> The security of the Libra Blockchain rests on the correct
> 
> 
> implementation of validators, Move programs, and the Move VM.
> 
> 
> Addressing these issues in Libra Core is a work in progress.

这部分内容已经基本总结完毕，尽管他们在Rust中编写了实现，这对性能和安全性来说似乎是一个良好的开端.


##  表现 Performance

> We anticipate the initial launch of Libra protocol to support 1,000
> 
> 
> payment transactions per second with a 10-second finality time between
> 
> 
> a transaction being submitted and committed.

由于只有100个左右的验证者，并且它们都相互直接连接的，所以10秒的块时间听起来是可行的.


最低节点要求：

* 40 Mbps网络连接
* 1个商品CPU
* 16 TB SSD
前面有一些关于保持验证人从头执行初始同步的能力，而不是信任来自其他验证人签名状态的参考文献.我预计，如果Libra得到充分使用，那么执行这样的同步将很快变得非常不切实际，因此，节点安全模型将高度依赖于信任验证者.


## 用Move实现Libra生态系统策略

> The [Libra coin] reserve is the key mechanism for achieving value
> 
> 
> preservation. Through the reserve, each coin is fully backed with a
> 
> 
> set of stable and liquid assets. The Libra coin contract allows the
> 
> 
> association to mint new coins when demand increases and destroy them
> 
> 
> when the demand contracts. The association does not set a monetary
> 
> 
> policy. It can only mint and burn coins in response to demand from
> 
> 
> authorized resellers. Users do not need to worry about the association
> 
> 
> introducing inflation into the system or debasing the currency: For
> 
> 
> new coins to be minted, there must be a commensurate fiat deposit in
> 
> 
> the reserve.

好的，但现在我们讨论的是网络外部的事件.如白皮书前面所述，网络无法执行使用网络状态外部数据输入的脚本.因此，上述代码片段中的“can”和“must”修饰语肯定是指网络并不知道的Libra Association政策或合同义务.

> The consensus algorithm relies on the validator-set management Move
> 
> 
> module to maintain the current set of validators and manage the
> 
> 
> allocation of votes among the validators. Initially, the Libra
> 
> 
> Blockchain only grants votes to Founding Members.

假设验证者对验证者集的更改进行投票，听起来这会导致与我们在股权证明系统中看到的类似问题——远程攻击.如果创始成员的密匙的重要阈值受到损害，攻击者是否可以从源头写入新的账户历史记录？如果是这样，其他节点会接受吗？目前尚不清楚共识协议是否允许重写旧状态还是仅仅允许追加状态.

> We plan to gradually transition to a proof-of-stake.

如果他们能解决尚未解决的问题.

**未解决的问题**

**如何进行管理？**

我们可以看到Libra Association是一个由成员组成的委员会，需要2/3的绝对多数通过才能做出改变的决策.他们是唯一有资格铸造或销毁Libra coin的人，但如果有足够的共识，他们可以做出任何他们想要的改变.

**是否需要AML / KYC？**

显然，协议级别不需要它，但Calibra钱包声明所有用户都将通过政府颁发的ID进行验证.听起来Calibra钱包将是在一段时间内唯一可用的钱包，所以目前还不清楚开发人员和用户是否可以在Libra网络上运行不遵守与Calibra相同标准的应用程序.


什么是低收费？什么是正常操作？什么是足够的容量？


CALIBRA钱包FAQ承诺低收费，但这似乎与在高负载时底层协议的操作相冲突.

> Transaction fees will be low-cost and transparent, especially if
> 
> 
> you’re sending money internationally. Calibra will cut fees to help
> 
> 
> people keep more of their money.

**Libra真的会对开发者开放吗？**

根据实现无许可共识的计划：

> The Libra Blockchain will be open to everyone—any consumer, developer,
> 
> 
> or business can use the Libra network, build products on top of it,
> 
> 
> and add value through their services. Open access ensures low barriers
> 
> 
> to entry and innovation and encourages healthy competition that
> 
> 
> benefits consumers.

我怀疑开发人员是否能够在这个平台上运行他们所想像的任何技术上有效的应用程序.我没有读到任何让我相信这个系统会抵制审查制度的内容，但只有时间会告诉我们答案！

点击“ [Libra Blockchain](https://developers.libra.org/docs/assets/papers/the-libra-blockchain.pdf)”可查看原文


_扫码关注京东云开发者社区，每天都有精彩行业信息哦！_






