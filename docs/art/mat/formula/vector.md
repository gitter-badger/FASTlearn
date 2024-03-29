# 平面向量

```mermaid
graph TB
0(实际背景) --> 1(向量的概念) --> 2(向量的运算及其几何意义) ---- 3.1(向量的加减运算和几何性质) & 3.2(向量的数乘运算及其几何意义) & 3.3(向量的数量积及其几何意义) ---- 4(平面向量基本定理及其坐标表示) --- 5(平面向量的应用)
```

## 平面向量的线性计算

### 平面向量的加法

`a+0=0+a=a`

$\boldsymbol{a+0=0+a=a}$

`a+b<=a+b`

$\boldsymbol{|a+b|≤|a|+|b|}$

`(a+b)+c=a+(b+c)`

$\boldsymbol{(a+b)+c=a+(b+c)}$

### 平面向量的减法

`-(-a)=a`

$\boldsymbol{-(-a)=a}$

`a+(-a)=(-a)+a=0`

$\boldsymbol{a+(-a)=(-a)+a=0}$

`a-b=a+(-b)`

$\boldsymbol{a-b=a+(-b)}$

### 平面向量的数乘

引入：

`(-a)+(-a)+(-a)=-3a`

$\boldsymbol{(-a)+(-a)+(-a)=-3a}$

`-3a=3a`

$\boldsymbol{|-3a|=3|a|}$

根据定义得：

`l(ma)=(lm)a`

$\boldsymbol{\lambda(\mu a)=(\lambda \mu)a}$

`(lm)a=la+lb`

$\boldsymbol{(\lambda+\mu)a=\lambda a+\lambda b}$

`l(a+b)=la+lb`

$\boldsymbol{\lambda(a+b)=\lambda a+\lambda b}$

`(-l)a=-(la)=l(-a)`

$\boldsymbol{(-\lambda)a=-(\lambda a)=\lambda (-a)}$

`l(a-b)=la-lb`

$\boldsymbol{\lambda(a-b)=\lambda a-\lambda b}$

概括得：`l(m1a+-m2b)=lm1a+-lm2b`

$\boldsymbol{\lambda({\mu_1}a±{\mu_2}b)=\lambda {\mu_1} a±\lambda {\mu_2} b}$

!!! Example
    $\boldsymbol{(2a+3b-c)-(3a-2b+c)=2a+3b-c-3a+2b-c=-a+5b-2c}$

若 $\boldsymbol{a}$ 与 $\boldsymbol{b}$ 共线，则：`b=la`

$b=\lambda a$

## 平面向量的数量积

引入：`w=fscost`

$W=\boldsymbol{|F||s|\cos \theta}$

设 $\boldsymbol{a}$ 与 $\boldsymbol{b}$ 的夹角为 $\theta$ ，则：`ab=abct`

$\boldsymbol{a·b=|a||b|\cos \theta}$

根据定义，得：

`ab=ba`

$\boldsymbol{a·b=b·a}$

`(la)b=l(ab)=a(lb)`

$\boldsymbol{(\lambda a)·b=\lambda(a·b)=a·(\lambda b)}$

`(a+b)c=a·c+bc`

$\boldsymbol{(a+b)·c=a·c+b·c}$

## 平面向量的坐标表示

如果 $\boldsymbol{e_1}$ ， $\boldsymbol{e_2}$ 是同一平面内两个不共线向量那么对于这一平面内的任意向量 $\boldsymbol{a}$ ，有且只有一个实数 $\lambda_1$ ，$\lambda_2$ ，使：`a=l1e1+l2e2`

$\boldsymbol{a=\lambda_1 e_1+\lambda_2 e_2}$

此时，{$\boldsymbol{e_1}$ ，$\boldsymbol{e_2}$} 是该平面所有向量的基底，任何一个向量都可以由同一个基底唯一表示，即 $\boldsymbol{a} = x\boldsymbol{e_1}+y\boldsymbol{e_2}$.

向量的和差的坐标等于其坐标的和差，即：`a-b=x1x2,y1y2`

$\boldsymbol{a-b} = (x_1 - x_2, y_1 - y_2)$

向量与实数的积的坐标等于原向量与实数的积的坐标，即：`la=lx,ly`

$\lambda \boldsymbol{a}=(\lambda x, \lambda y)$

两个向量的数量积等于它们对应坐标的乘积的和，即：`ab=x1x2+y1y2`

$\boldsymbol{a·b} = x_1 x_2 + y_1 y_2$

若 $\boldsymbol{a⊥b}$ ，则：`ab=0`

$\boldsymbol{a·b} = 0$

## 平面向量与三角形

引入：`f=g/(2cost/2)`

$\boldsymbol{|F_1|=}\frac{\boldsymbol{|G|}}{2\cos \frac {\theta} 2}$

### 余弦定理

$a^2=b^2+c^2-2bc\cos A$

推论：`cosa=(b2c2-a2)/(2bc)`

$\cos A = \frac {b^2+c^2-a^2} {2bc}$

### 正弦定理

`a/sina=a/sinb=c/sinc`

$\frac a {\sin A}=\frac b {\sin B}=\frac c {\sin C}=2R$

其中，$R$ 是三角形外接圆的半径。

在 $△ABC$ 中， $A>B\Leftrightarrow\sin{A}>\sin{B}\Leftrightarrow{a>b}$

### 面积公式

$S=\sqrt{p(p-a)(p-b)(p-c)}, p=\frac 1 2 (a+b+c)$

$S=\frac{1}2ah_a=\frac{1}2bh_b=\frac{1}2ch_c$

$S=\frac{1}2ab\sin{C}=\frac{1}2bc\sin{A}=\frac{1}2ca\sin{B}$
