// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pactus

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func TransactionClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Transaction"),
		Short: "Transaction service client",
		Long:  "Transaction service defines various RPC methods for interacting with\n transactions.",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_TransactionGetTransactionCommand(cfg),
		_TransactionCalculateFeeCommand(cfg),
		_TransactionBroadcastTransactionCommand(cfg),
		_TransactionGetRawTransferTransactionCommand(cfg),
		_TransactionGetRawBondTransactionCommand(cfg),
		_TransactionGetRawUnbondTransactionCommand(cfg),
		_TransactionGetRawWithdrawTransactionCommand(cfg),
	)
	return cmd
}

func _TransactionGetTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &GetTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetTransaction"),
		Short: "GetTransaction RPC client",
		Long:  "GetTransaction retrieves transaction details based on the provided request\n parameters.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "GetTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &GetTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Id, cfg.FlagNamer("Id"), "Transaction ID.")
	flag.EnumVar(cmd.PersistentFlags(), &req.Verbosity, cfg.FlagNamer("Verbosity"), "Verbosity level for transaction details.")

	return cmd
}

func _TransactionCalculateFeeCommand(cfg *client.Config) *cobra.Command {
	req := &CalculateFeeRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("CalculateFee"),
		Short: "CalculateFee RPC client",
		Long:  "CalculateFee calculates the transaction fee based on the specified amount\n and payload type.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "CalculateFee"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &CalculateFeeRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.CalculateFee(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Float64Var(&req.Amount, cfg.FlagNamer("Amount"), 0, "Transaction amount.")
	flag.EnumVar(cmd.PersistentFlags(), &req.PayloadType, cfg.FlagNamer("PayloadType"), "Type of transaction payload.")
	cmd.PersistentFlags().BoolVar(&req.FixedAmount, cfg.FlagNamer("FixedAmount"), false, "Indicates that amount should be fixed and includes the fee.")

	return cmd
}

func _TransactionBroadcastTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &BroadcastTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("BroadcastTransaction"),
		Short: "BroadcastTransaction RPC client",
		Long:  "BroadcastTransaction broadcasts a signed transaction to the network.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "BroadcastTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &BroadcastTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.BroadcastTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	flag.BytesBase64Var(cmd.PersistentFlags(), &req.SignedRawTransaction, cfg.FlagNamer("SignedRawTransaction"), "Signed raw transaction data.")

	return cmd
}

func _TransactionGetRawTransferTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &GetRawTransferTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetRawTransferTransaction"),
		Short: "GetRawTransferTransaction RPC client",
		Long:  "GetRawTransferTransaction retrieves raw details of a transfer transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "GetRawTransferTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &GetRawTransferTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetRawTransferTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Uint32Var(&req.LockTime, cfg.FlagNamer("LockTime"), 0, "Lock time for the transaction.\n If not explicitly set, it sets to the last block height.")
	cmd.PersistentFlags().StringVar(&req.Sender, cfg.FlagNamer("Sender"), "", "Sender's account address.")
	cmd.PersistentFlags().StringVar(&req.Receiver, cfg.FlagNamer("Receiver"), "", "Receiver's account address.")
	cmd.PersistentFlags().Float64Var(&req.Amount, cfg.FlagNamer("Amount"), 0, "Transfer amount.\n It should be greater than 0.")
	cmd.PersistentFlags().Float64Var(&req.Fee, cfg.FlagNamer("Fee"), 0, "Transaction fee.\n If not explicitly set, it is calculated based on the amount.")
	cmd.PersistentFlags().StringVar(&req.Memo, cfg.FlagNamer("Memo"), "", "Transaction memo.")

	return cmd
}

func _TransactionGetRawBondTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &GetRawBondTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetRawBondTransaction"),
		Short: "GetRawBondTransaction RPC client",
		Long:  "GetRawBondTransaction retrieves raw details of a bond transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "GetRawBondTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &GetRawBondTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetRawBondTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Uint32Var(&req.LockTime, cfg.FlagNamer("LockTime"), 0, "Lock time for the transaction.\n If not explicitly set, it sets to the last block height.")
	cmd.PersistentFlags().StringVar(&req.Sender, cfg.FlagNamer("Sender"), "", "Sender's account address.")
	cmd.PersistentFlags().StringVar(&req.Receiver, cfg.FlagNamer("Receiver"), "", "Receiver's validator address.")
	cmd.PersistentFlags().Int64Var(&req.Stake, cfg.FlagNamer("Stake"), 0, "Stake amount.\n It should be greater than 0.")
	cmd.PersistentFlags().StringVar(&req.PublicKey, cfg.FlagNamer("PublicKey"), "", "Public key of the validator.")
	cmd.PersistentFlags().Float64Var(&req.Fee, cfg.FlagNamer("Fee"), 0, "Transaction fee.\n If not explicitly set, it is calculated based on the stake.")
	cmd.PersistentFlags().StringVar(&req.Memo, cfg.FlagNamer("Memo"), "", "Transaction memo.")

	return cmd
}

func _TransactionGetRawUnbondTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &GetRawUnbondTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetRawUnbondTransaction"),
		Short: "GetRawUnbondTransaction RPC client",
		Long:  "GetRawUnbondTransaction retrieves raw details of an unbond transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "GetRawUnbondTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &GetRawUnbondTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetRawUnbondTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Uint32Var(&req.LockTime, cfg.FlagNamer("LockTime"), 0, "Lock time for the transaction.\n If not explicitly set, it sets to the last block height.")
	cmd.PersistentFlags().StringVar(&req.ValidatorAddress, cfg.FlagNamer("ValidatorAddress"), "", "Address of the validator to unbond from.")
	cmd.PersistentFlags().StringVar(&req.Memo, cfg.FlagNamer("Memo"), "", "Transaction memo.")

	return cmd
}

func _TransactionGetRawWithdrawTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &GetRawWithdrawTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetRawWithdrawTransaction"),
		Short: "GetRawWithdrawTransaction RPC client",
		Long:  "GetRawWithdrawTransaction retrieves raw details of a withdraw transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Transaction", "GetRawWithdrawTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTransactionClient(cc)
				v := &GetRawWithdrawTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetRawWithdrawTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Uint32Var(&req.LockTime, cfg.FlagNamer("LockTime"), 0, "Lock time for the transaction.\n If not explicitly set, it sets to the last block height.")
	cmd.PersistentFlags().StringVar(&req.ValidatorAddress, cfg.FlagNamer("ValidatorAddress"), "", "Address of the validator to withdraw from.")
	cmd.PersistentFlags().StringVar(&req.AccountAddress, cfg.FlagNamer("AccountAddress"), "", "Address of the account to withdraw to.")
	cmd.PersistentFlags().Float64Var(&req.Amount, cfg.FlagNamer("Amount"), 0, "Withdrawal amount.\n It should be greater than 0.")
	cmd.PersistentFlags().Float64Var(&req.Fee, cfg.FlagNamer("Fee"), 0, "Transaction fee.\n If not explicitly set, it is calculated based on the stake.")
	cmd.PersistentFlags().StringVar(&req.Memo, cfg.FlagNamer("Memo"), "", "Transaction memo.")

	return cmd
}
