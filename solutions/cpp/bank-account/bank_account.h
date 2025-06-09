#if !defined(BANK_ACCOUNT_H)
#define BANK_ACCOUNT_H

#include <stdexcept>
#include <mutex>

namespace Bankaccount {

class Bankaccount {

public:
    Bankaccount( ) { account_open=false; account_balance=0; };
    void open( );
    void close( );
    int  balance( );
    void deposit(  int deposit_amount );
    void withdraw( int withdrawal_amount );

private:
    bool   account_open;
    int    account_balance;
    std::mutex account_mutex;

};  // class Bankaccount

}  // namespace Bankaccount

#endif  // BANK_ACCOUNT_H

