# word2vec

- The underlying assumption of Word2Vec is that two words sharing similar contexts also share a similar meaning and consequently a similar vector representation from the model.

### one-hot encoding

- 产生大量冗余稀疏矩阵
- 单词间的关系没有体现

### word hashing

![word hashing](.\word hashing.png)



### word embedding

- 输入one-hot Vector
- Hidden Layer没有激活函数（现行单元）
- Output Layer 和Inpyt Layer 维度一样，用的是softmax函数
- 模型训练好以后，并不会用这个模型处理新的任务，而是使用参数。（eg.权重矩阵W）

***模型定义输入与输出***

有两种方法：

1. CBOW（Continuous Bag-of-words)
	\- 输入：某特征词的上下文的词向量
	\- 输出：这个特征词的词向量
2. Skip-Gram
	\- 输入：某个特征词的词向量
	\- 输出：这个特征词的上下文的词向量

CBOW一般适用于小型语料库；Skip-Gram可以很好地运用在大型语料库中。

`CBOW`模型训练图

![](https://www.pianshen.com/images/980/5b69ad9ff59c1ffa446dad43081a25ac.png)

1. 输入层：上下文单词的onehot. (假设单词向量空间dim为V，上下文单词个数为C)
2. 所有onehot分别乘以共享的输入权重矩阵W. (V*N矩阵，N为自己设定的数，初始化权重矩阵W)
3. 所得的向量 (因为是onehot所以为向量) 相加求平均作为隐层向量, size为1N.
4. 乘以输出权重矩阵W’ (NV)
5. 得到向量 (1*V) softmax函数处理得到V-dim概率分布
6. 概率最大的index所指示的单词为预测出的中间词（target word）与true label的onehot做比较，误差越小越好（根据误差更新权重矩阵）



