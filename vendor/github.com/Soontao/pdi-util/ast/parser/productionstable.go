// Code generated by gocc; DO NOT EDIT.

package parser

import ast "github.com/Soontao/pdi-util/ast/types"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : StatementList	<< ast.NewProgram(X[0], nil) >>`,
		Id:         "Program",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProgram(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Program : StatementList BusinessObjectDefination	<< ast.NewProgram(X[0], X[1]) >>`,
		Id:         "Program",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProgram(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `StatementList : Statement	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "StatementList",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `StatementList : StatementList Statement	<< ast.NewCommonList(X[0], X[1]) >>`,
		Id:         "StatementList",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Statement : ImportStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : VariableDeclarationStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : ForeachStmt	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : IfStmt	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : RaiseStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : ReturnStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : AssignmentStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `Statement : ExpressionStmt terminator	<< ast.NewStatement(X[0]) >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatement(X[0])
		},
	},
	ProdTabEntry{
		String: `ImportStmt : "import" Namespace	<< ast.NewImportDeclaration(X[1], nil) >>`,
		Id:         "ImportStmt",
		NTType:     4,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewImportDeclaration(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `ImportStmt : "import" Namespace keywordAs Identifier	<< ast.NewImportDeclaration(X[1], X[3]) >>`,
		Id:         "ImportStmt",
		NTType:     4,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewImportDeclaration(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `ReturnStmt : "return"	<<  >>`,
		Id:         "ReturnStmt",
		NTType:     5,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ReturnStmt : "return" Expression	<<  >>`,
		Id:         "ReturnStmt",
		NTType:     5,
		Index:      16,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `VariableDeclarationStmt : "var" Identifier	<< ast.NewVariableDeclarationStmt(X[1]), nil >>`,
		Id:         "VariableDeclarationStmt",
		NTType:     6,
		Index:      17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariableDeclarationStmt(X[1]), nil
		},
	},
	ProdTabEntry{
		String: `VariableDeclarationStmt : VariableDeclarationStmt "=" Expression	<< ast.GeneralAddProperty(X[0], "Value", X[2]), nil >>`,
		Id:         "VariableDeclarationStmt",
		NTType:     6,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Value", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `VariableDeclarationStmt : VariableDeclarationStmt ":" DataType	<< ast.GeneralAddProperty(X[0], "DataType", X[2]), nil >>`,
		Id:         "VariableDeclarationStmt",
		NTType:     6,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "DataType", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `AssignmentStmt : Expression "=" Expression	<<  >>`,
		Id:         "AssignmentStmt",
		NTType:     7,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ForeachStmt : "foreach" "(" "var" Identifier "in" Expression ")" StatementsBlock	<< ast.NewForeachStmt(X[3], X[5], X[7]), nil >>`,
		Id:         "ForeachStmt",
		NTType:     8,
		Index:      21,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewForeachStmt(X[3], X[5], X[7]), nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" "(" Expression ")" StatementsBlock	<< ast.NewIfStmt(X[2], X[4]), nil >>`,
		Id:         "IfStmt",
		NTType:     9,
		Index:      22,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfStmt(X[2], X[4]), nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : IfStmt "else" StatementsBlock	<< ast.GeneralAddProperty(X[0], "Else", X[2]), nil >>`,
		Id:         "IfStmt",
		NTType:     9,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Else", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : IfStmt "else" IfStmt	<< ast.GeneralAddProperty(X[0], "ElseOf", X[2]), nil >>`,
		Id:         "IfStmt",
		NTType:     9,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "ElseOf", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `RaiseStmt : "raise" Expression	<<  >>`,
		Id:         "RaiseStmt",
		NTType:     10,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ExpressionStmt : Expression	<<  >>`,
		Id:         "ExpressionStmt",
		NTType:     11,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StatementsBlock : "{" StatementList "}"	<< X[1], nil >>`,
		Id:         "StatementsBlock",
		NTType:     12,
		Index:      27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : keywordBusinessObject DataType "{" BOItemList "}"	<< ast.NewBODefination(nil, X[1], nil, X[3] ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      28,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(nil, X[1], nil, X[3] )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : keywordBusinessObject DataType "{" "}"	<< ast.NewBODefination(nil, X[1], nil, nil ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      29,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(nil, X[1], nil, nil )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : keywordBusinessObject DataType RaiseExpr "{" BOItemList "}"	<< ast.NewBODefination(nil, X[1], X[2], X[5] ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      30,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(nil, X[1], X[2], X[5] )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : AnnotationList keywordBusinessObject DataType "{" BOItemList "}"	<< ast.NewBODefination(X[0], X[2], nil, X[4] ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      31,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(X[0], X[2], nil, X[4] )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : AnnotationList keywordBusinessObject DataType "{" "}"	<< ast.NewBODefination(X[0], X[2], nil, nil ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      32,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(X[0], X[2], nil, nil )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : AnnotationList keywordBusinessObject DataType RaiseExpr "{" BOItemList "}"	<< ast.NewBODefination(X[0], X[2], X[3], X[5] ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      33,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(X[0], X[2], X[3], X[5] )
		},
	},
	ProdTabEntry{
		String: `BusinessObjectDefination : AnnotationList keywordBusinessObject DataType RaiseExpr "{" "}"	<< ast.NewBODefination(X[0], X[2], X[3], nil ) >>`,
		Id:         "BusinessObjectDefination",
		NTType:     13,
		Index:      34,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBODefination(X[0], X[2], X[3], nil )
		},
	},
	ProdTabEntry{
		String: `BOItemList : AnnotatedBOItem	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "BOItemList",
		NTType:     14,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `BOItemList : BOItemList AnnotatedBOItem	<< ast.NewCommonList(X[0], X[1]) >>`,
		Id:         "BOItemList",
		NTType:     14,
		Index:      36,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `AnnotatedBOItem : AnnotationList BOItem	<< ast.NewAnnotatedBOItem(X[0], X[1]) >>`,
		Id:         "AnnotatedBOItem",
		NTType:     15,
		Index:      37,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnnotatedBOItem(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `AnnotatedBOItem : BOItem	<< ast.NewAnnotatedBOItem(nil, X[0]) >>`,
		Id:         "AnnotatedBOItem",
		NTType:     15,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnnotatedBOItem(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `BusinessObjectNode : keywordNode Identifier "{" BOItemList "}"	<< ast.NewBusinessObjectNode(X[1], X[3]) >>`,
		Id:         "BusinessObjectNode",
		NTType:     16,
		Index:      39,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBusinessObjectNode(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `BusinessObjectNode : keywordNode Identifier "{" "}"	<< ast.NewBusinessObjectNode(X[1], nil) >>`,
		Id:         "BusinessObjectNode",
		NTType:     16,
		Index:      40,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBusinessObjectNode(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `BusinessObjectNode : keywordNode Identifier Multiplicity "{" BOItemList "}"	<< ast.NewBusinessObjectNode(X[1], X[4], X[2]) >>`,
		Id:         "BusinessObjectNode",
		NTType:     16,
		Index:      41,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBusinessObjectNode(X[1], X[4], X[2])
		},
	},
	ProdTabEntry{
		String: `BusinessObjectNode : keywordNode Identifier RaiseExpr "{" BOItemList "}"	<< ast.NewBusinessObjectNode(X[1], X[4], nil, X[2]) >>`,
		Id:         "BusinessObjectNode",
		NTType:     16,
		Index:      42,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBusinessObjectNode(X[1], X[4], nil, X[2])
		},
	},
	ProdTabEntry{
		String: `BusinessObjectNode : keywordNode Identifier Multiplicity RaiseExpr "{" BOItemList "}"	<< ast.NewBusinessObjectNode(X[1], X[5], X[2], X[3]) >>`,
		Id:         "BusinessObjectNode",
		NTType:     16,
		Index:      43,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBusinessObjectNode(X[1], X[5], X[2], X[3])
		},
	},
	ProdTabEntry{
		String: `BOItem : ElementItem	<< X[0], nil >>`,
		Id:         "BOItem",
		NTType:     17,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BOItem : AssociationItem	<< X[0], nil >>`,
		Id:         "BOItem",
		NTType:     17,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BOItem : MessageItem	<< X[0], nil >>`,
		Id:         "BOItem",
		NTType:     17,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BOItem : ActionItem	<< X[0], nil >>`,
		Id:         "BOItem",
		NTType:     17,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BOItem : BusinessObjectNode	<< X[0], nil >>`,
		Id:         "BOItem",
		NTType:     17,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageItem : "message" Identifier "text" Value ":" DataTypeList terminator	<< ast.NewMessageItem(X[1], X[3], X[5]), nil >>`,
		Id:         "MessageItem",
		NTType:     18,
		Index:      49,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMessageItem(X[1], X[3], X[5]), nil
		},
	},
	ProdTabEntry{
		String: `AssociationItem : "association" MultiplicityIdentifier	<< ast.NewAssociationItem(X[1], nil, nil, nil), nil >>`,
		Id:         "AssociationItem",
		NTType:     19,
		Index:      50,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAssociationItem(X[1], nil, nil, nil), nil
		},
	},
	ProdTabEntry{
		String: `AssociationItem : AssociationItem "to" DataType	<< ast.GeneralAddProperty(X[0], "Target", X[2]), nil >>`,
		Id:         "AssociationItem",
		NTType:     19,
		Index:      51,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Target", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `AssociationItem : AssociationItem "valuation" "(" Expression ")"	<< ast.GeneralAddProperty(X[0], "Valuation", X[3]), nil >>`,
		Id:         "AssociationItem",
		NTType:     19,
		Index:      52,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Valuation", X[3]), nil
		},
	},
	ProdTabEntry{
		String: `AssociationItem : AssociationItem "using" Identifier	<< ast.GeneralAddProperty(X[0], "Using", X[2]), nil >>`,
		Id:         "AssociationItem",
		NTType:     19,
		Index:      53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Using", X[2]), nil
		},
	},
	ProdTabEntry{
		String: `AssociationItem : AssociationItem terminator	<< X[0], nil >>`,
		Id:         "AssociationItem",
		NTType:     19,
		Index:      54,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ElementItem : "element" Identifier ":" DataType terminator	<< ast.NewElementItem(X[1], X[3], nil) >>`,
		Id:         "ElementItem",
		NTType:     20,
		Index:      55,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewElementItem(X[1], X[3], nil)
		},
	},
	ProdTabEntry{
		String: `ElementItem : "element" Identifier ":" DataType "=" Value terminator	<< ast.NewElementItem(X[1], X[3], X[5]) >>`,
		Id:         "ElementItem",
		NTType:     20,
		Index:      56,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewElementItem(X[1], X[3], X[5])
		},
	},
	ProdTabEntry{
		String: `ActionItem : keywordAction Identifier terminator	<< ast.NewActionItem(X[1], nil), nil >>`,
		Id:         "ActionItem",
		NTType:     21,
		Index:      57,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewActionItem(X[1], nil), nil
		},
	},
	ProdTabEntry{
		String: `ActionItem : keywordAction Identifier RaiseExpr terminator	<< ast.NewActionItem(X[1], X[2]), nil >>`,
		Id:         "ActionItem",
		NTType:     21,
		Index:      58,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewActionItem(X[1], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `ExpressionList : Expression	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "ExpressionList",
		NTType:     22,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `ExpressionList : ExpressionList "," Expression	<< ast.NewCommonList(X[0], X[2]) >>`,
		Id:         "ExpressionList",
		NTType:     22,
		Index:      60,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expression : Expression "||" Term1	<< ast.NewExpression("||", X[0], X[2]), nil >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      61,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("||", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Expression : Term1	<<  >>`,
		Id:         "Expression",
		NTType:     23,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term1 : Term1 "&&" Term2	<< ast.NewExpression("&&", X[0], X[2]), nil >>`,
		Id:         "Term1",
		NTType:     24,
		Index:      63,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("&&", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term1 : Term2	<<  >>`,
		Id:         "Term1",
		NTType:     24,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term2 : Term2 RelOp Term3	<< ast.NewExpression(X[1], X[0], X[2]), nil >>`,
		Id:         "Term2",
		NTType:     25,
		Index:      65,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression(X[1], X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term2 : Term3	<<  >>`,
		Id:         "Term2",
		NTType:     25,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term3 : Term3 "+" Term4	<< ast.NewExpression("+", X[0], X[2]), nil >>`,
		Id:         "Term3",
		NTType:     26,
		Index:      67,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("+", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term3 : Term4	<<  >>`,
		Id:         "Term3",
		NTType:     26,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term4 : Term4 "-" Term5	<< ast.NewExpression("-", X[0], X[2]), nil >>`,
		Id:         "Term4",
		NTType:     27,
		Index:      69,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("-", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term4 : Term5	<<  >>`,
		Id:         "Term4",
		NTType:     27,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term5 : Term5 "*" Term6	<< ast.NewExpression("*", X[0], X[2]), nil >>`,
		Id:         "Term5",
		NTType:     28,
		Index:      71,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("*", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term5 : Term6	<<  >>`,
		Id:         "Term5",
		NTType:     28,
		Index:      72,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term6 : Term6 "/" Term7	<< ast.NewExpression("/", X[0], X[2]), nil >>`,
		Id:         "Term6",
		NTType:     29,
		Index:      73,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("/", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term6 : Term7	<<  >>`,
		Id:         "Term6",
		NTType:     29,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term7 : Term7 "%!"(MISSING) Term8	<< ast.NewExpression("%", X[0], X[2]), nil >>`,
		Id:         "Term7",
		NTType:     30,
		Index:      75,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression("%", X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Term7 : Term8	<<  >>`,
		Id:         "Term7",
		NTType:     30,
		Index:      76,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term8 : "(" Expression ")"	<< X[1], nil >>`,
		Id:         "Term8",
		NTType:     31,
		Index:      77,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Term8 : UnaryExpr	<<  >>`,
		Id:         "Term8",
		NTType:     31,
		Index:      78,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryExpr : PrimaryExpr	<<  >>`,
		Id:         "UnaryExpr",
		NTType:     32,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryExpr : UnaryOp UnaryExpr	<<  >>`,
		Id:         "UnaryExpr",
		NTType:     32,
		Index:      80,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : Identifier	<< X[0], nil >>`,
		Id:         "PrimaryExpr",
		NTType:     33,
		Index:      81,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : Value	<< X[0], nil >>`,
		Id:         "PrimaryExpr",
		NTType:     33,
		Index:      82,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : Selector	<< X[0], nil >>`,
		Id:         "PrimaryExpr",
		NTType:     33,
		Index:      83,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : PrimaryExpr Arguments	<< ast.NewFuncCallExpr(X[0], X[1]), nil >>`,
		Id:         "PrimaryExpr",
		NTType:     33,
		Index:      84,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncCallExpr(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `UnaryOp : "+"	<<  >>`,
		Id:         "UnaryOp",
		NTType:     34,
		Index:      85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryOp : "-"	<<  >>`,
		Id:         "UnaryOp",
		NTType:     34,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryOp : "!"	<<  >>`,
		Id:         "UnaryOp",
		NTType:     34,
		Index:      87,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryOp : "*"	<<  >>`,
		Id:         "UnaryOp",
		NTType:     34,
		Index:      88,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryOp : "&"	<<  >>`,
		Id:         "UnaryOp",
		NTType:     34,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "=="	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "!="	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<="	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RelOp : "<"	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">="	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RelOp : ">"	<< ast.ConvertToString(X[0]), nil >>`,
		Id:         "RelOp",
		NTType:     35,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConvertToString(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `Arguments : "(" ")"	<<  >>`,
		Id:         "Arguments",
		NTType:     36,
		Index:      96,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Arguments : "(" ExpressionList ")"	<<  >>`,
		Id:         "Arguments",
		NTType:     36,
		Index:      97,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DataTypeList : DataType	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "DataTypeList",
		NTType:     37,
		Index:      98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `DataTypeList : DataTypeList "," DataType	<< ast.NewCommonList(X[0], X[2]) >>`,
		Id:         "DataTypeList",
		NTType:     37,
		Index:      99,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `DataType : Namespace ":" Identifier	<< ast.NewDataType(X[0], X[2]) >>`,
		Id:         "DataType",
		NTType:     38,
		Index:      100,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDataType(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `DataType : Identifier	<< ast.NewDataType(X[0], nil) >>`,
		Id:         "DataType",
		NTType:     38,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDataType(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `KeyValueList : KeyValue	<< ast.NewKeyValueList(X[0], nil) >>`,
		Id:         "KeyValueList",
		NTType:     39,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyValueList(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `KeyValueList : KeyValueList "," KeyValue	<< ast.NewKeyValueList(X[2], X[0]) >>`,
		Id:         "KeyValueList",
		NTType:     39,
		Index:      103,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyValueList(X[2], X[0])
		},
	},
	ProdTabEntry{
		String: `KeyValue : Identifier "=" Value	<< ast.NewKeyValuePair(X[0], X[2]) >>`,
		Id:         "KeyValue",
		NTType:     40,
		Index:      104,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKeyValuePair(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `NamespaceList : Namespace	<<  >>`,
		Id:         "NamespaceList",
		NTType:     41,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NamespaceList : Namespace "," NamespaceList	<<  >>`,
		Id:         "NamespaceList",
		NTType:     41,
		Index:      106,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AnnotationList : Annotation	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "AnnotationList",
		NTType:     42,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `AnnotationList : AnnotationList Annotation	<< ast.NewCommonList(X[0], X[1]) >>`,
		Id:         "AnnotationList",
		NTType:     42,
		Index:      108,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Namespace : Identifier	<< ast.NewNamespace(X[0], nil) >>`,
		Id:         "Namespace",
		NTType:     43,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNamespace(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Namespace : Identifier "." Namespace	<< ast.NewNamespace(X[0], X[2]) >>`,
		Id:         "Namespace",
		NTType:     43,
		Index:      110,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNamespace(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Selector : Identifier	<< ast.NewSelector(X[0], nil) >>`,
		Id:         "Selector",
		NTType:     44,
		Index:      111,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSelector(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Selector : Identifier "." Selector	<< ast.NewSelector(X[0], X[2]) >>`,
		Id:         "Selector",
		NTType:     44,
		Index:      112,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSelector(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Annotation : "[" Identifier "]"	<< ast.NewAnnotation(X[1], nil, nil) >>`,
		Id:         "Annotation",
		NTType:     45,
		Index:      113,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnnotation(X[1], nil, nil)
		},
	},
	ProdTabEntry{
		String: `Annotation : "[" Identifier "(" Value ")" "]"	<< ast.NewAnnotation(X[1], nil, X[3]) >>`,
		Id:         "Annotation",
		NTType:     45,
		Index:      114,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnnotation(X[1], nil, X[3])
		},
	},
	ProdTabEntry{
		String: `Annotation : "[" Identifier "(" Namespace ")" "]"	<< ast.NewAnnotation(X[1], X[3], nil) >>`,
		Id:         "Annotation",
		NTType:     45,
		Index:      115,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAnnotation(X[1], X[3], nil)
		},
	},
	ProdTabEntry{
		String: `RepeatTerminator : terminator RepeatTerminator	<<  >>`,
		Id:         "RepeatTerminator",
		NTType:     46,
		Index:      116,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RepeatTerminator : empty	<<  >>`,
		Id:         "RepeatTerminator",
		NTType:     46,
		Index:      117,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `RaiseExpr : keywordRaises IdentifierList	<< ast.NewRaiseExpr(X[1]) >>`,
		Id:         "RaiseExpr",
		NTType:     47,
		Index:      118,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRaiseExpr(X[1])
		},
	},
	ProdTabEntry{
		String: `IdentifierList : Identifier	<< ast.NewCommonList(nil, X[0]) >>`,
		Id:         "IdentifierList",
		NTType:     48,
		Index:      119,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `IdentifierList : IdentifierList "," Identifier	<< ast.NewCommonList(X[0], X[2]) >>`,
		Id:         "IdentifierList",
		NTType:     48,
		Index:      120,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCommonList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `MultiplicityIdentifier : Identifier	<< X[0], nil >>`,
		Id:         "MultiplicityIdentifier",
		NTType:     49,
		Index:      121,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MultiplicityIdentifier : Identifier Multiplicity	<< ast.GeneralAddProperty(X[0], "Multiplicity", X[1]), nil >>`,
		Id:         "MultiplicityIdentifier",
		NTType:     49,
		Index:      122,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GeneralAddProperty(X[0], "Multiplicity", X[1]), nil
		},
	},
	ProdTabEntry{
		String: `Identifier : identifier	<< ast.NewIdentifier(X[0]) >>`,
		Id:         "Identifier",
		NTType:     50,
		Index:      123,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIdentifier(X[0])
		},
	},
	ProdTabEntry{
		String: `Multiplicity : "[" Value "," Value "]"	<< ast.NewMultiplicity(X[1], X[3]) >>`,
		Id:         "Multiplicity",
		NTType:     51,
		Index:      124,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMultiplicity(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Multiplicity : "[" Value "," "n" "]"	<< ast.NewMultiplicity(X[1], "n") >>`,
		Id:         "Multiplicity",
		NTType:     51,
		Index:      125,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMultiplicity(X[1], "n")
		},
	},
	ProdTabEntry{
		String: `Value : floatLit	<< ast.NewNumberValue(X[0]) >>`,
		Id:         "Value",
		NTType:     52,
		Index:      126,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNumberValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : intLit	<< ast.NewNumberValue(X[0]) >>`,
		Id:         "Value",
		NTType:     52,
		Index:      127,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNumberValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : stringLit	<< ast.NewStringValue(X[0]) >>`,
		Id:         "Value",
		NTType:     52,
		Index:      128,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : boolLit	<< ast.NewBoolValue(X[0]) >>`,
		Id:         "Value",
		NTType:     52,
		Index:      129,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : ComplexValue	<< X[0], nil >>`,
		Id:         "Value",
		NTType:     52,
		Index:      130,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ComplexValue : "{" KeyValueList "}"	<< ast.NewComplexValue(X[1]) >>`,
		Id:         "ComplexValue",
		NTType:     53,
		Index:      131,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComplexValue(X[1])
		},
	},
}
